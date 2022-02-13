/*
Copyright © 2020 Richard Nixon

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

package types

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// ESSnapshot represents an ES snapshotsitory from from the _cat/snapshots URI
type ESSnapshot struct {
	Id               string `json:"id"`
	Status           string `json:"status"`
	StartEpoch       string `json:"start_epoch"`
	StartTime        string `json:"start_time"`
	EndEpoch         string `json:"end_epoch"`
	EndTime          string `json:"end_time"`
	Duration         string `json:"duration"`
	Indices          string `json:"indices"`
	SuccessfulShards string `json:"successful_shards"`
	FailedShards     string `json:"failed_shards"`
	TotalShards      string `json:"total_shards"`
}

// Unmarshal parses Json into an ESSnapshot object (singular)
func (essnapshot *ESSnapshot) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, essnapshot)
}

// ESSnapshots represents a collection of ES snapshots from the _cat/snapshots/<repo> URI
type ESSnapshots []ESSnapshot

// Unmarshal parses Json into an ESNodes object (array)
func (essnapshots *ESSnapshots) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, essnapshots)
}

// PrintTable displays data from an ESNodes object
func (essnapshots *ESSnapshots) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Start Time", "End Time", "Duration", "Successful Shards", "Failed Shards"})
	for _, snap := range *essnapshots {
		table.Append([]string{snap.Id, snap.Status, snap.StartTime, snap.EndTime, snap.Duration, snap.SuccessfulShards, snap.FailedShards})
	}
	table.Render()
}

// PrintJSON displays data from an ESSnapshots object
func (essnapshots *ESSnapshots) PrintJSON() {
	jsonBytes, _ := json.MarshalIndent(essnapshots, "", "  ")
	fmt.Println(string(jsonBytes))
}
