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
 * Copyright 2019 gRPC authors.
 *
 */

// Package wrr contains the interface and common implementations of wrr
// algorithms.
package wrr

// WRR defines an interface that implements weighted round robin.
type WRR interface {
	// Add adds an item with weight to the WRR set.
	//
	// Add and Next need to be thread safe.
	Add(item any, weight int64)
	// Next returns the next picked item.
	//
	// Add and Next need to be thread safe.
	Next() any
}
