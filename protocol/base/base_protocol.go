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

package base

import (
	"sync"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

// Protocol is the interface that wraps the basic Export, Refer and Destroy method.
type Protocol interface {
	// Export method is to export service for remote invocation
	Export(invoker Invoker) Exporter
	// Refer method is to refer a remote service
	Refer(url *common.URL) Invoker
	// Destroy method will destroy all invokers and exporters, so it only is called once.
	Destroy()
}

// BaseProtocol is default protocol implement.
type BaseProtocol struct {
	exporterMap *sync.Map
	invokers    []Invoker
}

// NewBaseProtocol creates a new BaseProtocol
func NewBaseProtocol() BaseProtocol {
	return BaseProtocol{
		exporterMap: new(sync.Map),
	}
}

// SetExporterMap set @exporter with @key to local memory.
func (bp *BaseProtocol) SetExporterMap(key string, exporter Exporter) {
	bp.exporterMap.Store(key, exporter)
}

// ExporterMap gets exporter map.
func (bp *BaseProtocol) ExporterMap() *sync.Map {
	return bp.exporterMap
}

// SetInvokers sets invoker into local memory
func (bp *BaseProtocol) SetInvokers(invoker Invoker) {
	bp.invokers = append(bp.invokers, invoker)
}

// Invokers gets all invokers
func (bp *BaseProtocol) Invokers() []Invoker {
	return bp.invokers
}

// Export is default export implement.
func (bp *BaseProtocol) Export(invoker Invoker) Exporter {
	return NewBaseExporter("base", invoker, bp.exporterMap)
}

// Refer is default refer implement.
func (bp *BaseProtocol) Refer(url *common.URL) Invoker {
	return NewBaseInvoker(url)
}

// Destroy will destroy all invoker and exporter, so it only is called once.
func (bp *BaseProtocol) Destroy() {
	// destroy invokers
	for _, invoker := range bp.invokers {
		if invoker != nil {
			invoker.Destroy()
		}
	}
	bp.invokers = []Invoker{}

	// un export exporters
	bp.exporterMap.Range(func(key, exporter any) bool {
		if exporter != nil {
			exporter.(Exporter).UnExport()
		} else {
			bp.exporterMap.Delete(key)
		}
		return true
	})
}
