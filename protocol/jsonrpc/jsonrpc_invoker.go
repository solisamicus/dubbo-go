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

package jsonrpc

import (
	"context"
)

import (
	"github.com/dubbogo/gost/log/logger"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	"dubbo.apache.org/dubbo-go/v3/protocol/result"
)

// JsonrpcInvoker is JSON RPC invoker
type JsonrpcInvoker struct {
	base.BaseInvoker
	client *HTTPClient
}

// NewJsonrpcInvoker creates JSON RPC invoker with @url and @client
func NewJsonrpcInvoker(url *common.URL, client *HTTPClient) *JsonrpcInvoker {
	return &JsonrpcInvoker{
		BaseInvoker: *base.NewBaseInvoker(url),
		client:      client,
	}
}

// Invoke the JSON RPC invocation and return result.
func (ji *JsonrpcInvoker) Invoke(ctx context.Context, inv base.Invocation) result.Result {
	var result result.RPCResult

	rpcInv := inv.(*invocation.RPCInvocation)
	url := ji.GetURL()
	req := ji.client.NewRequest(url, rpcInv.MethodName(), rpcInv.Arguments())
	ctxNew := context.WithValue(ctx, constant.DubboGoCtxKey, map[string]string{
		"X-Proxy-ID": "dubbogo",
		"X-Services": url.Path,
		"X-Method":   rpcInv.MethodName(),
	})
	result.Err = ji.client.Call(ctxNew, url, req, rpcInv.Reply())
	if result.Err == nil {
		result.Rest = rpcInv.Reply()
	}
	logger.Debugf("result.Err: %v, result.Rest: %v", result.Err, result.Rest)

	return &result
}
