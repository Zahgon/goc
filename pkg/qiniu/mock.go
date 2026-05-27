/*
 Copyright 2020 Qiniu Cloud (qiniu.com)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package qiniu

import (
	"github.com/julienschmidt/httprouter"
)

// MockQiniuServer simulate qiniu cloud for testing
func MockQiniuServer(config *Config) (client *QnClient, router *httprouter.Router, serverURL string, teardown func()) {
	_ = "STUB: not implemented"
	// router is the HTTP request multiplexer used with the test server.
	return nil, nil, "", nil
}

// server is a test HTTP server used to provide mock API responses.

// MockRouterAPI mocks qiniu /v2/list API.
// You need to provide a expected profile content.
// count controls the mocks qiniu server to error before 'count' times request.
func MockRouterAPI(router *httprouter.Router, profile string, count int) {
	_ = "STUB: not implemented"

	// mock rsf /v2/list
	return
}

// mock io get statusJSON file

// mock io get remote coverage profile

// MockRouterListAllAPI mocks qiniu /list API.
// count controls the mocks qiniu server to error before 'count' times request.
func MockRouterListAllAPI(router *httprouter.Router, count int) {
	_ = "STUB: not implemented"

	// mock rsf /v2/list
	return
}

// MockPrivateDomainUrl mocks bucket domain /key, /timeout, /retry API.
// count controls the mocks qiniu server to error before 'count' times request.
func MockPrivateDomainUrl(router *httprouter.Router, count int) { _ = "STUB: not implemented"; return }
