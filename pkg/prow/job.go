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

package prow

import (
	"os"

	"github.com/qiniu/goc/pkg/cover"
	"github.com/qiniu/goc/pkg/github"
	"github.com/qiniu/goc/pkg/qiniu"
)

// IProwAction defines the normal action in prow system
type IProwAction interface {
	Fetch(BuildID, name string) []byte
	RunPresubmit() error
	RunPostsubmit() error
	RunPeriodic() error
}

// Job is a prowjob in prow
type Job struct {
	JobName                string
	Org                    string
	RepoName               string
	PRNumStr               string
	BuildId                string //prow job build number
	PostSubmitJob          string
	PostSubmitCoverProfile string
	CovThreshold           int
	LocalProfilePath       string
	QiniuClient            qiniu.Client
	LocalArtifacts         qiniu.Artifacts
	GithubComment          github.PrComment
	FullDiff               bool
}

// Fetch the file from cloud
func (j *Job) Fetch(BuildID, name string) []byte {
	_ = "STUB: not implemented"

	// RunPresubmit run a presubmit job
	return nil
}

func (j *Job) RunPresubmit() error {
	_ = "STUB: not implemented"
	// step1: get local profile cov
	return nil
}

//step2: find the remote healthy cover profile from qiniu bucket

// step3: get github pull request changed files' name and calculate diff cov between local and remote profile

// step4: generate changed file html coverage

// step5: post comment to github

// RunPostsubmit run a postsubmit job
func (j *Job) RunPostsubmit() error {
	_ = "STUB: not implemented"

	// RunPeriodic run a periodic job
	return nil
}

func (j *Job) RunPeriodic() error {
	_ = "STUB: not implemented"

	// trim github filename to profile format:
	//
	//	src/qiniu.com/kodo/io/io/io_svr.go -> qiniu.com/kodo/io/io/io_svr.go
	return nil
}

func trimGhFileToProfile(ghFiles []string) (pFiles []string) {
	_ = "STUB: not implemented"
	// TODO: need compatible other situation
	return nil
}

// WriteChangedCov filter local profile with changed files and save to j.LocalArtifacts.ChangedProfileName
func (j *Job) WriteChangedCov(changedFiles []string) error { _ = "STUB: not implemented"; return nil }

// writeLine writes a line in the given file, if the file pointer is not nil
func writeLine(file *os.File, content string) { _ = "STUB: not implemented"; return }

// JobPrefixOnQiniu generates the prefix string of the job on qiniu
func (j *Job) JobPrefixOnQiniu() string { _ = "STUB: not implemented"; return "" }

// HtmlProfile generates the name of the profile html file
func (j *Job) HtmlProfile() string { _ = "STUB: not implemented"; return "" }

// SetDeltaCovLinks set DeltaCovLinks to the job
func (j *Job) SetDeltaCovLinks(c cover.DeltaCovList) { _ = "STUB: not implemented"; return }

// CreateChangedCovHtml create changed file related coverage html base on the local artifact
func (j *Job) CreateChangedCovHtml() error { _ = "STUB: not implemented"; return nil }

func getFilesAndCovList(fullDiff bool, prComment github.PrComment, localP, baseP cover.CoverageList) (changedFiles []string, deltaCovList cover.DeltaCovList, err error) {
	_ = "STUB: not implemented"

	// get github pull request changed files' name
	return nil, *new(cover.DeltaCovList), nil
}

// calculate diff cov between local and remote profile
