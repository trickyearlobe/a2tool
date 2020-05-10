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

// ESNode represents an ES node from the _cat/nodes URI
type ESNode struct {
	IP          string `json:"ip"`
	HeapPercent string `json:"heap.percent"`
	RAMPercent  string `json:"ram.percent"`
	CPU         string `json:"cpu"`
	Load1m      string `json:"load_1m"`
	Load5m      string `json:"load_5m"`
	Load15m     string `json:"load_15m"`
	Role        string `json:"node.role"`
	Master      string `json:"master"`
	Name        string `json:"name"`
}

// Unmarshal parses Json into an ESNode object (singular)
func (esnode *ESNode) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esnode)
}

// ESNodes represents a collection of ES nodes from the _cat/nodes URI
type ESNodes []ESNode

// Unmarshal parses Json into an ESNodes object (array)
func (esnodes *ESNodes) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esnodes)
}

// PrintTable displays data from an ESNodes object
func (esnodes *ESNodes) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "IP", "Role", "Master", "CPU", "RAM%", "Heap%", "1min Ave", "5min Ave", "15min Ave"})
	for _, node := range *esnodes {
		table.Append([]string{node.Name, node.IP, node.Role, node.Master, node.CPU, node.RAMPercent, node.HeapPercent, node.Load1m, node.Load5m, node.Load15m})
	}
	table.Render()
}

// PrintJSON displays data from en ESNodes object
func (esnodes *ESNodes) PrintJSON() {
	jsonBytes, _ := json.MarshalIndent(esnodes, "", "  ")
	fmt.Println(string(jsonBytes))
}
