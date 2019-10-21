// Copyright 2019 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package common

import (
	"time"

	"github.com/gabriel-samfira/coriolis-logger/logging"
	"github.com/gabriel-samfira/coriolis-logger/params"
	"github.com/gabriel-samfira/coriolis-logger/worker"
	client "github.com/influxdata/influxdb1-client/v2"
)

type DataStore interface {
	worker.SimpleWorker

	Write(logMsg logging.LogMessage) error
	Rotate(olderThan time.Time) error
	ResultReader(p params.QueryParams) Reader
	List() ([]map[string]string, error)
	Query(q client.Query) (*client.ChunkedResponse, error)
}

type Reader interface {
	ReadNext() ([]byte, error)
}
