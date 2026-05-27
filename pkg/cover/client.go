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

package cover

import (
	"io"
	"net/http"
)

// Action provides methods to contact with the covered service under test
type Action interface {
	Profile(param ProfileParam) ([]byte, error)
	Clear(param ProfileParam) ([]byte, error)
	Remove(param ProfileParam) ([]byte, error)
	InitSystem() ([]byte, error)
	ListServices() ([]byte, error)
	RegisterService(svr ServiceUnderTest) ([]byte, error)
}

const (
	//CoverInitSystemAPI prepare a new round of testing
	CoverInitSystemAPI = "/v1/cover/init"
	//CoverProfileAPI is provided by the covered service to get profiles
	CoverProfileAPI = "/v1/cover/profile"
	//CoverProfileClearAPI is provided by the covered service to clear profiles
	CoverProfileClearAPI = "/v1/cover/clear"
	//CoverServicesListAPI list all the registered services
	CoverServicesListAPI = "/v1/cover/list"
	//CoverRegisterServiceAPI register a service into service center
	CoverRegisterServiceAPI = "/v1/cover/register"
	//CoverServicesRemoveAPI remove one services from the service center
	CoverServicesRemoveAPI = "/v1/cover/remove"
)

type client struct {
	Host   string
	client *http.Client
}

// NewWorker creates a worker to contact with service
func NewWorker(host string) Action { _ = "STUB: not implemented"; return *new(Action) }

func (c *client) RegisterService(srv ServiceUnderTest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) ListServices() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *client) Profile(param ProfileParam) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the json.Marshal function can return two types of errors: UnsupportedTypeError or UnsupportedValueError
// so no need to check here

func (c *client) Clear(param ProfileParam) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the json.Marshal function can return two types of errors: UnsupportedTypeError or UnsupportedValueError
// so no need to check here

func (c *client) Remove(param ProfileParam) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the json.Marshal function can return two types of errors: UnsupportedTypeError or UnsupportedValueError
// so no need to check here

func (c *client) InitSystem() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *client) do(method, url, contentType string, body io.Reader) (*http.Response, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func isNetworkError(err error) bool { _ = "STUB: not implemented"; return false }
