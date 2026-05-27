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

package build

import (
	"os"

	"github.com/qiniu/goc/pkg/cover"
)

func (b *Build) cpLegacyProject() { _ = "STUB: not implemented"; return }

// Skip if already copied

// only cp dependency in root(current gopath),
// skip deps in other GOPATHs
func (b *Build) cpDepPackages(pkg *cover.Package, visited map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// Check if copied

// Skip if already copied

// Check if we can found in the root gopath

func (b *Build) cpNonStandardLegacy() { _ = "STUB: not implemented"; return }

// skipCopy skip copy .git dir and irregular files
func skipCopy(src string, info os.FileInfo) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
