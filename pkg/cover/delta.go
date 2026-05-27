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

// DeltaCov contains the info of a delta coverage
type DeltaCov struct {
	FileName    string
	BasePer     string
	NewPer      string
	DeltaPer    string
	LineCovLink string
}

// DeltaCovList is the list of DeltaCov
type DeltaCovList []DeltaCov

// GetFullDeltaCov get full delta coverage between new and base profile
func GetFullDeltaCov(newList CoverageList, baseList CoverageList) (delta DeltaCovList) {
	_ = "STUB: not implemented"
	return *new(DeltaCovList)
}

//if the file not in base profile, set None

//if the file not in new profile, set None

// GetDeltaCov get two profile diff cov
func GetDeltaCov(newList CoverageList, baseList CoverageList) (delta DeltaCovList) {
	_ = "STUB: not implemented"
	return *new(DeltaCovList)
}

// GetChFileDeltaCov get two profile diff cov of changed files
func GetChFileDeltaCov(newList CoverageList, baseList CoverageList, changedFiles []string) (list DeltaCovList) {
	_ = "STUB: not implemented"
	return *new(DeltaCovList)
}

// Delta calculate two coverage delta
func Delta(new Coverage, base Coverage) float32 { _ = "STUB: not implemented"; return 0 }

// TotalDelta calculate two coverage delta
func TotalDelta(new CoverageList, base CoverageList) float32 { _ = "STUB: not implemented"; return 0 }

// Map returns maps the file name to its DeltaCov for faster retrieval & membership check
func (d DeltaCovList) Map() map[string]DeltaCov { _ = "STUB: not implemented"; return nil }

// Sort sort DeltaCovList with filenames
func (d DeltaCovList) Sort() { _ = "STUB: not implemented"; return }

// Name returns the file name
func (c *DeltaCov) Name() string {
	_ = "STUB: not implemented"

	// GetLineCovLink get the LineCovLink of the DeltaCov
	return ""
}

func (c *DeltaCov) GetLineCovLink() string { _ = "STUB: not implemented"; return "" }

// SetLineCovLink set LineCovLink of the DeltaCov
func (c *DeltaCov) SetLineCovLink(link string) { _ = "STUB: not implemented"; return }
