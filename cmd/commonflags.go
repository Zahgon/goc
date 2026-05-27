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

package cmd

import (
	"github.com/spf13/pflag"
)

var (
	target            string
	center            string
	agentPort         AgentPort
	debugGoc          bool
	debugInCISyncFile string
	buildFlags        string
	singleton         bool

	goRunExecFlag  string
	goRunArguments string
)

var coverMode = CoverMode{
	mode: "count",
}

// addBasicFlags adds a
func addBasicFlags(cmdset *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// bind to viper

func addCommonFlags(cmdset *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// bind to viper

func addBuildFlags(cmdset *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// bind to viper

func addRunFlags(cmdset *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// bind to viper

// CoverMode represents the covermode when doing cover for source code
type CoverMode struct {
	mode string
}

func (m *CoverMode) String() string {
	_ = "STUB: not implemented"

	// Set sets the value to the CoverMode struct, use 'count' as default if v is empty
	return ""
}

func (m *CoverMode) Set(v string) error { _ = "STUB: not implemented"; return nil }

// Type returns the type of CoverMode
func (m *CoverMode) Type() string {
	_ = "STUB: not implemented"

	// AgentPort is the struct to do agentPort check
	return ""
}

type AgentPort struct {
	port string
}

func (agent *AgentPort) String() string {
	_ = "STUB: not implemented"

	// Set sets the value to the AgentPort struct
	return ""
}

func (agent *AgentPort) Set(v string) error { _ = "STUB: not implemented"; return nil }

// Type returns the type of AgentPort
func (agent *AgentPort) Type() string { _ = "STUB: not implemented"; return "" }
