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

// MvProjectsToTmp moves the projects into a temporary directory
func (b *Build) MvProjectsToTmp() error { _ = "STUB: not implemented"; return nil }

// fix #14: unable to build project not in GOPATH in legacy mode
// this kind of project does not have a pkg.Root value
// go 1.11, 1.12 has no pkg.Root,
// so add b.IsMod == false as secondary judgement

func (b *Build) mvProjectsToTmp() error { _ = "STUB: not implemented"; return nil }

// Delete previous tmp folder and its content

// Create a new tmp folder and a new importpath for storing cover variables

// traverse pkg list to get project meta info

// we should get corresponding working directory in temporary directory

// issue #14
// if b.Root == "", then the project is non-standard project
// known cases:
// 1. a legacy project, but not in any GOPATH, will cause the b.Root == ""

// go 1.11, 1.12 has no Build.Root

// tmpFolderName uses the first six characters of the input path's SHA256 checksum
// as the suffix.
func tmpFolderName(path string) string { _ = "STUB: not implemented"; return "" }

// tmpPackageName uses the first six characters of the input path's SHA256 checksum
// as the suffix.
func tmpPackageName(path string) string { _ = "STUB: not implemented"; return "" }

// traversePkgsList travse the Build.Pkgs list
// return Build.IsMod, tell if the project is a mod project
// return Build.Root:
// 1. the project root if it is a mod project,
// 2. current GOPATH if it is a legacy project,
// 3. some non-standard project, which Build.IsMod == false, Build.Root == nil
func (b *Build) traversePkgsList() (isMod bool, root string, err error) {
	_ = "STUB: not implemented"
	return false,

		// get root
		"", nil
}

// getTmpwd get the corresponding working directory in the temporary working directory
// and store it in the Build.tmpWorkdingDir
func (b *Build) getTmpwd() (string, error) { _ = "STUB: not implemented"; return "", nil }

// b.TmpWorkingDir = filepath.Join(b.TmpDir, path[len(parentPath):])

func (b *Build) findWhereToInstall() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Clean clears up the temporary workspace
func (b *Build) Clean() error { _ = "STUB: not implemented"; return nil }
