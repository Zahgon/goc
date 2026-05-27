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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Do coverage profile diff analysis, it can also work with prow and post comments to github pull request if needed",
	Example: `	# Diff two local coverage profile and display
	goc diff --new-profile=<xxxx> --base-profile=<xxxx> 

	# Diff local coverage profile with the remote one in prow job using default qiniu-credential
	goc diff --prow-postsubmit-job=<xxx> --new-profile=<xxx> 

	# Calculate and display full diff coverage between new-profile and base-profile, not concerned github changed files
	goc diff --prow-postsubmit-job=<xxx> --new-profile=<xxx> --full-diff=true 

	# Diff local coverage profile with the remote one in prow job
	goc diff --prow-postsubmit-job=<xxx> --prow-remote-profile-name=<xxx> 
    		 --qiniu-credential=<xxx> --new-profile=<xxxx> 

	# Diff coverage profile with the remote one in prow job, and post comments to github PR
	goc diff --prow-postsubmit-job=<xxx> --prow-profile=<xxx> 
    		 --github-token=<xxx> --github-user=<xxx> --github-comment-prefix=<xxx> 
    		 --qiniu-credential=<xxx> --coverage-threshold-percentage=<xxx> --new-profile=<xxxx> 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if baseProfile != "" {
			doDiffForLocalProfiles(cmd, args)
		} else if prowPostSubmitJob != "" {
			doDiffUnderProw(cmd, args)
		} else {
			logrus.Fatalf("either base-profile or prow-postsubmit-job must be provided")
		}
	},
}

var (
	newProfile        string
	baseProfile       string
	coverageThreshold int

	prowPostSubmitJob string
	prowProfile       string

	githubToken         string
	githubUser          string
	githubCommentPrefix string

	qiniuCredential string

	robotName string
	fullDiff  bool
)

func init() {
	diffCmd.Flags().StringVarP(&newProfile, "new-profile", "n", "", "local profile which works as the target to analysis")
	diffCmd.MarkFlagRequired("new-profile")
	diffCmd.Flags().StringVarP(&baseProfile, "base-profile", "b", "", "another local profile which works as baseline to compare with the target")
	diffCmd.Flags().IntVarP(&coverageThreshold, "coverage-threshold-percentage", "", 0, "coverage threshold percentage")
	diffCmd.Flags().StringVarP(&prowPostSubmitJob, "prow-postsubmit-job", "", "", "prow postsubmit job which used to find the base profile")
	diffCmd.Flags().StringVarP(&prowProfile, "prow-remote-profile-name", "", "filtered.cov", "the name of profile in prow postsubmit job, which used as the base profile to compare")
	diffCmd.Flags().StringVarP(&githubToken, "github-token", "", "/etc/github/oauth", "path to token to access github repo")
	diffCmd.Flags().StringVarP(&githubUser, "github-user", "", "", "github user name when comments in github")
	diffCmd.Flags().StringVarP(&githubCommentPrefix, "github-comment-prefix", "", "", "specific comment flag you provided")
	diffCmd.Flags().StringVarP(&qiniuCredential, "qiniu-credential", "", "/etc/qiniuconfig/qiniu.json", "path to credential file to access qiniu cloud")
	diffCmd.Flags().StringVarP(&robotName, "robot-name", "", "qiniu-bot", "github user name for coverage robot")
	diffCmd.Flags().BoolVarP(&fullDiff, "full-diff", "", false, "when set true,calculate and display full diff coverage between new-profile and base-profile")

	rootCmd.AddCommand(diffCmd)
}

// goc diff --new-profile=./new.cov --base-profile=./base.cov
// +------------------------------------------------------+---------------+--------------+--------+
// |                         File                         | Base Coverage | New Coverage | Delta  |
// +------------------------------------------------------+---------------+--------------+--------+
// | qiniu.com/kodo/bd/pfd/pfdstg/cursor/mgr.go           |     53.5%     |    50.5%     | -3.0%  |
// | qiniu.com/kodo/bd/pfd/pfdstg/svr/getstripe.go        |     0.5%      |     0.0%     | -0.5%  |
// | Total                                                |     35.7%     |    35.7%     | -0.0%  |
// +------------------------------------------------------+---------------+--------------+--------+
func doDiffForLocalProfiles(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

//calculate diff file cov and display

func doDiffUnderProw(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }
