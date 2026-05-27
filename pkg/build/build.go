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
	"github.com/qiniu/goc/pkg/cover"
)

// Build is to describe the building/installing process of a goc build/install
type Build struct {
	Pkgs          map[string]*cover.Package // Pkg list parsed from "go list -json ./..." command
	NewGOPATH     string                    // the new GOPATH
	OriGOPATH     string                    // the original GOPATH
	WorkingDir    string                    // the working directory
	TmpDir        string                    // the temporary directory to build the project
	TmpWorkingDir string                    // the working directory in the temporary directory, which is corresponding to the current directory in the project directory
	IsMod         bool                      // determine whether it is a Mod project
	Root          string
	// go 1.11, go 1.12 has no Root
	// Project Root:
	// 1. legacy, root == GOPATH
	// 2. mod, root == go.mod Dir
	ModRoot     string // path for go.mod
	ModRootPath string // import path for the whole project
	Target      string // the binary name that go build generate
	// keep compatible with go commands:
	// go run [build flags] [-exec xprog] package [arguments...]
	// go build [-o output] [-i] [build flags] [packages]
	// go install [-i] [build flags] [packages]
	BuildFlags     string // Build flags
	Packages       string // Packages that needs to build
	GoRunExecFlag  string // for the -exec flags in go run command
	GoRunArguments string // for the '[arguments]' parameters in go run command

	OneMainPackage           bool   // whether this build is a go build or go install? true: build, false: install
	GlobalCoverVarImportPath string // Importpath for storing cover variables
	GlobalCoverVarFilePath   string // Importpath for storing cover variables
}

// NewBuild creates a Build struct which can build from goc temporary directory,
// and generate binary in current working directory
func NewBuild(buildflags string, args []string, workingDir string, outputDir string) (*Build, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildflags = buildflags + " -o " + outputDir

// Build calls 'go build' tool to do building
func (b *Build) Build() error { _ = "STUB: not implemented"; return nil }

// new -o will overwrite  previous ones

// Change to temp GOPATH for go install command

// determineOutputDir, as we only allow . as package name,
// the binary name is always same as the directory name of current directory
func (b *Build) determineOutputDir(outputDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// fix #43

// fix #43
// use target name from `go list -json ./...` of the main module

// validatePackageForBuild only allow . as package name
func (b *Build) validatePackageForBuild() bool { _ = "STUB: not implemented"; return false }

func checkParameters(args []string, workingDir string) error { _ = "STUB: not implemented"; return nil }
