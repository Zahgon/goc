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
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge [files...]",
	Short: "Merge multiple coherent Go coverage files into a single file.",
	Long: `merge will merge multiple Go coverage files into a single coverage file.
merge requires that the files are 'coherent', meaning that if they both contain references to the
same paths, then the contents of those source files were identical for the binary that generated
each file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		runMerge(args, outputMergeProfile)
	},
}

var outputMergeProfile string

func init() {
	mergeCmd.Flags().StringVarP(&outputMergeProfile, "output", "o", "mergeprofile.cov", "output file")

	rootCmd.AddCommand(mergeCmd)
}

func runMerge(args []string, output string) { _ = "STUB: not implemented"; return }
