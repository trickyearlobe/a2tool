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

// ESRepo represents an ES repository from from the _cat/repositories URI
type ESRepo struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// Unmarshal parses Json into an ESRepo object (singular)
func (esrepo *ESRepo) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esrepo)
}

// ESRepos represents a collection of ES repos from the _cat/repositories URI
type ESRepos []ESRepo

// Unmarshal parses Json into an ESNodes object (array)
func (esrepos *ESRepos) Unmarshal(jsonBytes []byte) {
	json.Unmarshal(jsonBytes, esrepos)
}

// PrintTable displays data from an ESNodes object
func (esrepos *ESRepos) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Type"})
	for _, repo := range *esrepos {
		table.Append([]string{repo.Id, repo.Type})
	}
	table.Render()
}

// PrintJSON displays data from an ESRepos object
func (esrepos *ESRepos) PrintJSON() {
	jsonBytes, _ := json.MarshalIndent(esrepos, "", "  ")
	fmt.Println(string(jsonBytes))
}
