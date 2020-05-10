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

// ESShard represents an ES shard from the _cat/shards URI
type ESShard struct {
	Index  string `json:"index"`
	Shard  string `json:"shard"`
	PriRep string `json:"prirep"`
	State  string `json:"state"`
	Docs   string `json:"docs"`
	Store  string `json:"store"`
	IP     string `json:"ip"`
	Node   string `json:"node"`
}

// Unmarshal parses Json into an ESNode object (singular)
func (esshard *ESShard) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esshard)
}

// ESShard represents a collection of ES shard from the _cat/shards URI
type ESShards []ESShard

// Unmarshal parses Json into an ESNodes object (array)
func (esshards *ESShards) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esshards)
}

// PrintTable displays data from an ESShards object
func (esshards *ESShards) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	goodCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiGreenColor}
	badCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiRedColor}
	goodRow := []tablewriter.Colors{goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell}
	badRow := []tablewriter.Colors{badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell}
	table.SetHeader([]string{"Index", "Shard", "Node", "IP", "Pri/Rep", "State", "Docs", "Store"})
	for _, shard := range *esshards {
		if shard.State == "STARTED" {
			table.Rich([]string{shard.Index, shard.Shard, shard.Node, shard.IP, shard.PriRep, shard.State, shard.Docs, shard.Store}, goodRow)
		} else {
			table.Rich([]string{shard.Index, shard.Shard, shard.Node, shard.IP, shard.PriRep, shard.State, shard.Docs, shard.Store}, badRow)
		}
	}
	table.Render()
}

// PrintJSON displays data from en ESShards object
func (esshards *ESShards) PrintJSON() {
	jsonBytes, _ := json.MarshalIndent(esshards, "", "  ")
	fmt.Println(string(jsonBytes))
}
