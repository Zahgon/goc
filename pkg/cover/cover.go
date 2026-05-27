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
	"errors"
	"io"
	"os/exec"
	"time"
)

var (
	// ErrCoverPkgFailed represents the error that fails to inject the package
	ErrCoverPkgFailed = errors.New("fail to inject code to project")
	// ErrCoverListFailed represents the error that fails to list package dependencies
	ErrCoverListFailed = errors.New("fail to list package dependencies")
)

// TestCover is a collection of all counters
type TestCover struct {
	Mode                     string
	AgentPort                string
	Center                   string // cover profile host center
	Singleton                bool
	MainPkgCover             *PackageCover
	DepsCover                []*PackageCover
	CacheCover               map[string]*PackageCover
	GlobalCoverVarImportPath string
}

// PackageCover holds all the generate coverage variables of a package
type PackageCover struct {
	Package *Package
	Vars    map[string]*FileVar
}

// FileVar holds the name of the generated coverage variables targeting the named file.
type FileVar struct {
	File string
	Var  string
}

// Package map a package output by go list
// this is subset of package struct in: https://github.com/golang/go/blob/master/src/cmd/go/internal/load/pkg.go#L58
type Package struct {
	Dir        string `json:"Dir"`        // directory containing package sources
	ImportPath string `json:"ImportPath"` // import path of package in dir
	Name       string `json:"Name"`       // package name
	Target     string `json:",omitempty"` // installed target for this package (may be executable)
	Root       string `json:",omitempty"` // Go root, Go path dir, or module root dir containing this package

	Module   *ModulePublic `json:",omitempty"`         // info about package's module, if any
	Goroot   bool          `json:"Goroot,omitempty"`   // is this package in the Go root?
	Standard bool          `json:"Standard,omitempty"` // is this package part of the standard Go library?
	DepOnly  bool          `json:"DepOnly,omitempty"`  // package is only a dependency, not explicitly listed

	// Source files
	GoFiles  []string `json:"GoFiles,omitempty"`  // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
	CgoFiles []string `json:"CgoFiles,omitempty"` // .go source files that import "C"

	// Dependency information
	Deps      []string          `json:"Deps,omitempty"` // all (recursively) imported dependencies
	Imports   []string          `json:",omitempty"`     // import paths used by this package
	ImportMap map[string]string `json:",omitempty"`     // map from source import to ImportPath (identity entries omitted)

	// Error information
	Incomplete bool            `json:"Incomplete,omitempty"` // this package or a dependency has an error
	Error      *PackageError   `json:"Error,omitempty"`      // error loading package
	DepsErrors []*PackageError `json:"DepsErrors,omitempty"` // errors loading dependencies
}

// ModulePublic represents the package info of a module
type ModulePublic struct {
	Path      string        `json:",omitempty"` // module path
	Version   string        `json:",omitempty"` // module version
	Versions  []string      `json:",omitempty"` // available module versions
	Replace   *ModulePublic `json:",omitempty"` // replaced by this module
	Time      *time.Time    `json:",omitempty"` // time version was created
	Update    *ModulePublic `json:",omitempty"` // available update (with -u)
	Main      bool          `json:",omitempty"` // is this the main module?
	Indirect  bool          `json:",omitempty"` // module is only indirectly needed by main module
	Dir       string        `json:",omitempty"` // directory holding local copy of files, if any
	GoMod     string        `json:",omitempty"` // path to go.mod file describing module, if any
	GoVersion string        `json:",omitempty"` // go version used in module
	Error     *ModuleError  `json:",omitempty"` // error loading module
}

// ModuleError represents the error loading module
type ModuleError struct {
	Err string // error text
}

// PackageError is the error info for a package when list failed
type PackageError struct {
	ImportStack []string // shortest path from package named on command line to this one
	Pos         string   // position of error (if present, file:line:col)
	Err         string   // the error itself
}

// CoverBuildInfo retreives some info from build
type CoverInfo struct {
	Target                   string
	GoPath                   string
	IsMod                    bool
	ModRootPath              string
	GlobalCoverVarImportPath string // path for the injected global cover var file
	OneMainPackage           bool
	Args                     string
	Mode                     string
	AgentPort                string
	Center                   string
	Singleton                bool
}

// Execute inject cover variables for all the .go files in the target folder
func Execute(coverInfo *CoverInfo) error { _ = "STUB: not implemented"; return nil }

// oneMainPackage := coverInfo.OneMainPackage

// var seenCache = make(map[string]*PackageCover)

// inject the main package

// new a testcover for this service

// handle its dependency
// var internalPkgCache = make(map[string][]*PackageCover)

//only focus package neither standard Go library nor dependency library

// inject Http Cover APIs

// ListPackages list all packages under specific via go list command
// The argument newgopath is if you need to go list in a different GOPATH
func ListPackages(dir string, args string, newgopath string) (map[string]*Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for _, err := range pkg.DepsErrors {
// 	log.Fatalf("dependency package list failed, err: %v", err)
// }

// AddCounters is different from official go tool cover
// 1. only inject covervar++ into source file
// 2. no declarartions for these covervars
// 3. return the declarations as string
func AddCounters(pkg *Package, mode string, globalCoverVarImportPath string) (*PackageCover, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func isDirExist(path string) bool { _ = "STUB: not implemented"; return false }

// Refer: https://github.com/golang/go/blob/master/src/cmd/go/internal/load/pkg.go#L1334:6
// hasInternalPath looks for the final "internal" path element in the given import path.
// If there isn't one, hasInternalPath returns ok=false.
// Otherwise, hasInternalPath returns ok=true and the index of the "internal".
func hasInternalPath(path string) bool {
	_ = "STUB: not implemented"
	// Three cases, depending on internal at start/end of string or not.
	// The order matters: we must return the index of the final element,
	// because the final one produces the most restrictive requirement
	// on the importer.
	return false
}

func getInternalParent(path string) string { _ = "STUB: not implemented"; return "" }

func buildCoverCmd(file string, coverVar *FileVar, pkg *Package, mode, newgopath string) *exec.Cmd {
	_ = "STUB: not implemented"
	// to construct: go tool cover -mode=atomic -o dest src (note: dest==src)
	return nil
}

// declareCoverVars attaches the required cover variables names
// to the files, to be used when annotating the files.
func declareCoverVars(p *Package) map[string]*FileVar { _ = "STUB: not implemented"; return nil }

// We create the cover counters as new top-level variables in the package.
// We need to avoid collisions with user variables (GoCover_0 is unlikely but still)
// and more importantly with dot imports of other covered packages,
// so we append 12 hex digits from the SHA-256 of the import path.
// The point is only to avoid accidents, not to defeat users determined to
// break things.

// These names appear in the cmd/cover HTML interface.

// These names appear in the cmd/cover HTML interface.

func declareCacheVars(in *PackageCover) map[string]*FileVar { _ = "STUB: not implemented"; return nil }

func cacheInternalCover(in *PackageCover) *PackageCover { _ = "STUB: not implemented"; return nil }

func addCacheCover(pkg *Package, in *PackageCover) *PackageCover {
	_ = "STUB: not implemented"
	return nil
}

// CoverageList is a collection and summary over multiple file Coverage objects
type CoverageList []Coverage

// Coverage stores test coverage summary data for one file
type Coverage struct {
	FileName      string
	NCoveredStmts int
	NAllStmts     int
	LineCovLink   string
}

type codeBlock struct {
	fileName      string // the file the code block is in
	numStatements int    // number of statements in the code block
	coverageCount int    // number of times the block is covered
}

// CovList converts profile to CoverageList struct
func CovList(f io.Reader) (g CoverageList, err error) {
	_ = "STUB: not implemented"
	return *new(CoverageList), nil
}

// discard first line

// ReadFileToCoverList coverts profile file to CoverageList struct
func ReadFileToCoverList(path string) (g CoverageList, err error) {
	_ = "STUB: not implemented"
	return *new(CoverageList), nil
}

// NewCoverageList return empty CoverageList
func NewCoverageList() CoverageList { _ = "STUB: not implemented"; return *new(CoverageList) }

func newCoverage(name string) *Coverage { _ = "STUB: not implemented"; return nil }

// convert a line in profile file to a codeBlock struct
func toBlock(line string) (res *codeBlock, err error) { _ = "STUB: not implemented"; return nil, nil }

// add blk Coverage to file group Coverage
func (blk *codeBlock) addToGroupCov(g *CoverageList) { _ = "STUB: not implemented"; return }

// when a new file name is processed

func (g CoverageList) size() int { _ = "STUB: not implemented"; return 0 }

func (g CoverageList) lastElement() *Coverage { _ = "STUB: not implemented"; return nil }

func (g *CoverageList) append(c *Coverage) { _ = "STUB: not implemented"; return }

// Sort sorts CoverageList with filenames
func (g CoverageList) Sort() { _ = "STUB: not implemented"; return }

// TotalPercentage returns the total percentage of coverage
func (g CoverageList) TotalPercentage() string { _ = "STUB: not implemented"; return "" }

// TotalRatio returns the total ratio of covered statements
func (g CoverageList) TotalRatio() (ratio float32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Map returns maps the file name to its coverage for faster retrieval
// & membership check
func (g CoverageList) Map() map[string]Coverage { _ = "STUB: not implemented"; return nil }

// Name returns the file name
func (c *Coverage) Name() string {
	_ = "STUB: not implemented"

	// Percentage returns the percentage of statements covered
	return ""
}

func (c *Coverage) Percentage() string { _ = "STUB: not implemented"; return "" }

// Ratio calculates the ratio of statements in a profile
func (c *Coverage) Ratio() (ratio float32, err error) { _ = "STUB: not implemented"; return 0, nil }

// PercentStr converts a fraction number to percentage string representation
func PercentStr(f float32) string { _ = "STUB: not implemented"; return "" }
