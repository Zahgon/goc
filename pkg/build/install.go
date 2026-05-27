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

// NewInstall creates a Build struct which can install from goc temporary directory
func NewInstall(buildflags string, args []string, workingDir string) (*Build, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Install use the 'go install' tool to install packages
func (b *Build) Install() error { _ = "STUB: not implemented"; return nil }

// ignore the err

// Change the temp GOBIN, to force binary install to original place

// Change to temp GOPATH for go install command

func (b *Build) validatePackageForInstall() bool { _ = "STUB: not implemented"; return false }
