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

// ESIndex represents an ES node index from the _cat/indices URI
type ESIndex struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	UUID         string `json:"uuid"`
	Primary      string `json:"pri"`
	Rep          string `json:"rep"`
	DocCount     string `json:"docs.count"`
	DocsDeleted  string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	PriStoreSize string `json:"pri.store.size"`
}

// Unmarshal parses Json into an ESIndex object (singular)
func (esindex *ESIndex) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esindex)
}

// ESIndices represents a collection of ESIndex from the _cat/indices URI
type ESIndices []ESIndex

// Unmarshal parses Json into an ESIndices object (singular)
func (esindices *ESIndices) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esindices)
}

// PrintTable displays data from an ESIndices object
func (esindices *ESIndices) PrintTable() {
	goodCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiGreenColor}
	badCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiRedColor}
	goodRow := []tablewriter.Colors{goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell}
	badRow := []tablewriter.Colors{badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell, badCell}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Health", "Status", "Docs", "Deleted", "LocalSize", "PriSize", "Pri", "Rep", "UUID"})
	for _, index := range *esindices {
		if index.Status == "open" {
			table.Rich([]string{index.Index, index.Health, index.Status, index.DocCount, index.DocsDeleted, index.StoreSize, index.PriStoreSize, index.Primary, index.Rep, index.UUID}, goodRow)
		} else {
			table.Rich([]string{index.Index, index.Health, index.Status, index.DocCount, index.DocsDeleted, index.StoreSize, index.PriStoreSize, index.Primary, index.Rep, index.UUID}, badRow)
		}
	}
	table.Render()
}

// PrintJSON displays data from en ESIndices object
func (esindices *ESIndices) PrintJSON() {
	jsonBytes, _ := json.MarshalIndent(esindices, "", "  ")
	fmt.Println(string(jsonBytes))
}
