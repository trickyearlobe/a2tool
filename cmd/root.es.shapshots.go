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
	"io/ioutil"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/trickyearlobe/a2tool/types"
)

func esGet(uri string) []byte {
	esURL := fmt.Sprintf("%s://%s:%s/%s", esScheme, esHost, esPort, uri)
	resp, err := http.Get(esURL)
	errorExit(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorExit(err)

	return body
}

// allocationCmd represents the allocation command
var snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Show data about ES snapshots",
	Run: func(cmd *cobra.Command, args []string) {

		var repos types.ESRepos
		var snapshots types.ESSnapshots
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Snapshot", "Repo", "Type", "Status", "Indices", "Start Time", "End Time", "Duration", "Successful Shards", "Failed Shards"})
		goodCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiGreenColor}
		badCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiRedColor}
		goodRow := []tablewriter.Colors{goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell}
		badRow := []tablewriter.Colors{badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell}

		// Snapshots live in repositories, so get the repository list
		repos.Unmarshal(esGet("_cat/repositories?format=json"))

		// Iterate the repos looking for snapshots
		for _, repo := range repos {
			snapshots.Unmarshal(esGet(fmt.Sprintf("_cat/snapshots/%s?format=json", repo.Id)))

			// Add snapshots to the table
			for _, snap := range snapshots {
				if snap.Status == "SUCCESS" {
					table.Rich([]string{snap.Id, repo.Id, repo.Type, snap.Status, snap.Indices, snap.StartTime, snap.EndTime, snap.Duration, snap.SuccessfulShards, snap.FailedShards}, goodRow)
				} else {
					table.Rich([]string{snap.Id, repo.Id, repo.Type, snap.Status, snap.Indices, snap.StartTime, snap.EndTime, snap.Duration, snap.SuccessfulShards, snap.FailedShards}, badRow)
				}
			}
		}
		table.SetAutoMergeCellsByColumnIndex([]int{0})
		table.Render()
	},
}

func init() {
	esCmd.AddCommand(snapshotsCmd)
}
