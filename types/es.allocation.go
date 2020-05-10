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
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// ESAllocation represents an ES node disk allocation from the _cat/allocation URI
type ESAllocation struct {
	Node        string `json:"node"`
	IP          string `json:"ip"`
	Host        string `json:"host"`
	Shards      string `json:"shards"`
	DiskIndices string `json:"disk.indices"`
	DiskAvail   string `json:"disk.avail"`
	DiskUsed    string `json:"disk.used"`
	DiskTotal   string `json:"disk.total"`
	DiskPercent string `json:"disk.percent"`
}

// Unmarshal parses Json into an ESNode object (singular)
func (esallocation *ESAllocation) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esallocation)
}

// ESNodes represents a collection of ES allocations from the _cat/allocations URI
type ESAllocations []ESAllocation

// Unmarshal parses Json into an ESAllocations object (array)
func (esallocations *ESAllocations) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esallocations)
}

// PrintTable displays data from an ESAllocations object
func (esallocations *ESAllocations) PrintTable() {
	goodCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiGreenColor}
	badCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiRedColor}
	goodRow := []tablewriter.Colors{goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell}
	badRow := []tablewriter.Colors{badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Node", "IP", "Host", "Shards", "Indices", "DiskAvail", "DiskUsed", "DiskTotal", "DiskPercent"})
	for _, alloc := range *esallocations {
		diskPercent, _ := strconv.Atoi(alloc.DiskPercent)
		if diskPercent < 71 {
			table.Rich([]string{alloc.Node, alloc.IP, alloc.Host, alloc.Shards, alloc.DiskIndices, alloc.DiskAvail, alloc.DiskUsed, alloc.DiskTotal, alloc.DiskPercent}, goodRow)
		} else {
			table.Rich([]string{alloc.Node, alloc.IP, alloc.Host, alloc.Shards, alloc.DiskIndices, alloc.DiskAvail, alloc.DiskUsed, alloc.DiskTotal, alloc.DiskPercent}, badRow)
		}
	}
	table.Render()
}

// PrintJSON displays data from en ESAllocation object
func (esallocations *ESAllocations) PrintJSON() {
	jsonBytes, _ := json.MarshalIndent(esallocations, "", "  ")
	fmt.Println(string(jsonBytes))
}
