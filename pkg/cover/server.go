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

	"github.com/gin-gonic/gin"
	"golang.org/x/tools/cover"
)

// LogFile a file to save log.
const LogFile = "goc.log"

type server struct {
	PersistenceFile string
	IPRevise        bool // whether to do ip revise during registering
	Store           Store
}

// NewFileBasedServer new a file based server with persistenceFile
func NewFileBasedServer(persistenceFile string) (*server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMemoryBasedServer new a memory based server without persistenceFile
func NewMemoryBasedServer() *server { _ = "STUB: not implemented"; return nil }

// Run starts coverage host center
func (s *server) Run(port string) { _ = "STUB: not implemented"; return }

// both log to stdout and file by default

// Router init goc server engine
func (s *server) Route(w io.Writer) *gin.Engine { _ = "STUB: not implemented"; return nil }

// api to show the registered services

// ServiceUnderTest is a entry under being tested
type ServiceUnderTest struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Address  string `form:"address" json:"address" binding:"required"`
	IPRevise string `form:"ip_revise" json:"ip_revise" binding:"-"` // whether to do ip revise during registering
}

// ProfileParam is param of profile API
type ProfileParam struct {
	Force             bool     `form:"force" json:"force"`
	Service           []string `form:"service" json:"service"`
	Address           []string `form:"address" json:"address"`
	CoverFilePatterns []string `form:"coverfile" json:"coverfile"`
	SkipFilePatterns  []string `form:"skipfile" json:"skipfile"`
}

// listServices list all the registered services
func (s *server) listServices(c *gin.Context) { _ = "STUB: not implemented"; return }

func (s *server) registerService(c *gin.Context) { _ = "STUB: not implemented"; return }

// valid scenario, keep going

// Prefer user's decision first.

// only for IPV4
// refer: https://github.com/qiniu/goc/issues/177

// profile API examples:
// POST /v1/cover/profile
// { "force": "true", "service":["a","b"], "address":["c","d"],"coverfile":["e","f"] }
func (s *server) profile(c *gin.Context) { _ = "STUB: not implemented"; return }

// filterProfile filters profiles of the packages matching the coverFile pattern
func filterProfile(coverFile []string, profiles []*cover.Profile) ([]*cover.Profile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no need to check again for the file

// skipProfile skips profiles of the packages matching the skipFile pattern
func skipProfile(skipFile []string, profiles []*cover.Profile) ([]*cover.Profile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no need to check again for the file

func (s *server) clear(c *gin.Context) { _ = "STUB: not implemented"; return }

func (s *server) initSystem(c *gin.Context) { _ = "STUB: not implemented"; return }

func (s *server) removeServices(c *gin.Context) { _ = "STUB: not implemented"; return }

func convertProfile(p []byte) ([]*cover.Profile, error) {
	_ = "STUB: not implemented"
	// Annoyingly, ParseProfiles only accepts a filename, so we have to write the bytes to disk
	// so it can read them back.
	// We could probably also just give it /dev/stdin, but that'll break on Windows.
	return nil, nil
}

func contains(arr []string, str string) bool { _ = "STUB: not implemented"; return false }

// filterAddrInfo filter address list by given service and address list
func filterAddrInfo(serviceList, addressList []string, force bool, allInfos map[string][]string) (filterAddrList []ServiceUnderTest, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add matched services to map

// jump to match the next service

// Add matched addresses to map

// Return all services when all param is nil
