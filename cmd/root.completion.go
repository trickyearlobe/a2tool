/*
Copyright © 2020 Richard Nixon <richard.nixon@btinternet.com>

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
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion script (requires bash-completions package)",

	Run: func(cmd *cobra.Command, args []string) {
		var completionsPathOutput string

		// If /etc/bash_completion.d exists, use that, otherwise use pkg-config
		stat, err := os.Stat("/etc/bash_completion.d")
		if err == nil {
			completionsPathOutput = "/etc/bash_completion.d"
			fmt.Printf("Found completions path at %s with stat %v\n", completionsPathOutput, stat)
		} else {
			completionsPathCmd := exec.Command("bash", "-c", "pkg-config --variable=completionsdir bash-completion")
			output, cmdErr := completionsPathCmd.CombinedOutput()
			completionsPathOutput = string(output)
			if cmdErr != nil {
				errorExit(fmt.Errorf("failed to dynamically locate completions path with `pkg-config --variable=completionsdir bash-completion`"))
			}
		}

		completionsPath := fmt.Sprintf("%s/a2tool", strings.TrimSuffix(completionsPathOutput, "\n"))
		completionFile, fileErr := os.Create(completionsPath)
		errorExit(fileErr)

		defer completionFile.Close()
		rootCmd.GenBashCompletion(completionFile)
		fmt.Printf("Completions written to %s\nPlease reload your session to activate changes\n", completionsPath)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
