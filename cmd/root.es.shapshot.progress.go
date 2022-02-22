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

	"github.com/spf13/cobra"
	"github.com/trickyearlobe/a2tool/types"
)

var flagSnapshotStatusDetail string

var esSnapshotProgressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Show data about running ES snapshots",
	Run: func(cmd *cobra.Command, args []string) {

		var status types.ESSnapshotStatus
		status.Unmarshal(esGet("_snapshot/_status"))

		switch flagSnapshotStatusDetail {
		case "repo":
			status.PrintRepoTable()
		case "index":
			status.PrintIndexTable()
		case "shard":
			status.PrintShardTable()
		default:
			fmt.Printf("--detail-level '%v' not supported\n", flagSnapshotStatusDetail)
		}
	},
}

func init() {
	esSnapshotCmd.AddCommand(esSnapshotProgressCmd)
	esSnapshotProgressCmd.PersistentFlags().StringVar(&flagSnapshotStatusDetail, "detail-level", "repo", "Level [repo index shard]")
}
