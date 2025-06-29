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

package rest

import (
	"sync"
)

import (
	"github.com/dubbogo/gost/log/logger"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
)

// nolint
type RestExporter struct {
	base.BaseExporter
}

// NewRestExporter returns a RestExporter
func NewRestExporter(key string, invoker base.Invoker, exporterMap *sync.Map) *RestExporter {
	return &RestExporter{
		BaseExporter: *base.NewBaseExporter(key, invoker, exporterMap),
	}
}

// UnExport unexport the RestExporter
func (re *RestExporter) UnExport() {
	interfaceName := re.GetInvoker().GetURL().GetParam(constant.InterfaceKey, "")
	err := common.ServiceMap.UnRegister(interfaceName, REST, re.GetInvoker().GetURL().ServiceKey())
	if err != nil {
		logger.Errorf("[RestExporter.UnExport] error: %v", err)
	}
	re.BaseExporter.UnExport()
}
