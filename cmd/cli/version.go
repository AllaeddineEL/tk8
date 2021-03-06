// Copyright © 2018 The TK8 Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// GITCOMMIT will hold the commit SHA to be used in the version command.
	GITCOMMIT string
	// VERSION will hold the version number to be used in the version command.
	VERSION = "dev-build"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of TK8",
	Long:  `All software has versions. This is TK8's`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Version:" + VERSION + " (Build: " + GITCOMMIT + ")")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
