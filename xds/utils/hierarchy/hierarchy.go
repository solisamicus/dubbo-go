/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 *
 * Copyright 2020 gRPC authors.
 *
 */

// Package hierarchy contains functions to set and get hierarchy string from
// addresses.
//
// This package is experimental.
package hierarchy

import (
	"google.golang.org/grpc/resolver"
)

type pathKeyType string

const pathKey = pathKeyType("grpc.internal.address.hierarchical_path")

type pathValue []string

func (p pathValue) Equal(o any) bool {
	op, ok := o.(pathValue)
	if !ok {
		return false
	}
	if len(op) != len(p) {
		return false
	}
	for i, v := range p {
		if v != op[i] {
			return false
		}
	}
	return true
}

// Get returns the hierarchical path of addr.
func Get(addr resolver.Address) []string {
	attrs := addr.BalancerAttributes
	if attrs == nil {
		return nil
	}
	path, _ := attrs.Value(pathKey).(pathValue)
	return ([]string)(path)
}

// Set overrides the hierarchical path in addr with path.
func Set(addr resolver.Address, path []string) resolver.Address {
	addr.BalancerAttributes = addr.BalancerAttributes.WithValue(pathKey, pathValue(path))
	return addr
}

// Group splits a slice of addresses into groups based on
// the first hierarchy path. The first hierarchy path will be removed from the
// result.
//
// Input:
// [
//
//	{addr0, path: [p0, wt0]}
//	{addr1, path: [p0, wt1]}
//	{addr2, path: [p1, wt2]}
//	{addr3, path: [p1, wt3]}
//
// ]
//
// Addresses will be split into p0/p1, and the p0/p1 will be removed from the
// path.
//
// Output:
//
//	{
//	  p0: [
//	    {addr0, path: [wt0]},
//	    {addr1, path: [wt1]},
//	  ],
//	  p1: [
//	    {addr2, path: [wt2]},
//	    {addr3, path: [wt3]},
//	  ],
//	}
//
// If hierarchical path is not set, or has no path in it, the address is
// dropped.
func Group(addrs []resolver.Address) map[string][]resolver.Address {
	ret := make(map[string][]resolver.Address)
	for _, addr := range addrs {
		oldPath := Get(addr)
		if len(oldPath) == 0 {
			continue
		}
		curPath := oldPath[0]
		newPath := oldPath[1:]
		newAddr := Set(addr, newPath)
		ret[curPath] = append(ret[curPath], newAddr)
	}
	return ret
}
