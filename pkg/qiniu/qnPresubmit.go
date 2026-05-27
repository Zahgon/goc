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
	"os"
)

const (
	//statusJSON is the JSON file that stores build success info
	statusJSON = "finished.json"

	// ArtifactsDirName is the name of directory defined in prow to store test artifacts
	ArtifactsDirName = "artifacts"

	//PostSubmitCoverProfile represents the default output coverage file generated in prow environment
	PostSubmitCoverProfile = "filtered.cov"

	//ChangedProfileName represents the default changed coverage profile based on files changed in Pull Request
	ChangedProfileName = "changed-file-profile.cov"
)

// sortBuilds converts all build from str to int and sorts all builds in descending order and
// returns the sorted slice
func sortBuilds(strBuilds []string) []int { _ = "STUB: not implemented"; return nil }

type finishedStatus struct {
	Timestamp int
	Passed    bool
}

func isBuildSucceeded(jsonText []byte) bool { _ = "STUB: not implemented"; return false }

// FindBaseProfileFromQiniu finds the coverage profile file from the latest healthy build
// stored in given gcs directory
func FindBaseProfileFromQiniu(qc Client, prowJobName, covProfileName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Artifacts is the interface of the rule to store test artifacts in prow
type Artifacts interface {
	ProfilePath() string
	CreateChangedProfile() *os.File
	GetChangedProfileName() string
}

// ProfileArtifacts presents the rule to store test artifacts in prow
type ProfileArtifacts struct {
	Directory          string
	ProfileName        string
	ChangedProfileName string // create temporary to save changed file related coverage profile
}

// ProfilePath returns a full path for profile
func (a *ProfileArtifacts) ProfilePath() string { _ = "STUB: not implemented"; return "" }

// CreateChangedProfile creates a profile in order to store the most related files based on Github Pull Request
func (a *ProfileArtifacts) CreateChangedProfile() *os.File { _ = "STUB: not implemented"; return nil }

// GetChangedProfileName get ChangedProfileName of the ProfileArtifacts
func (a *ProfileArtifacts) GetChangedProfileName() string { _ = "STUB: not implemented"; return "" }
