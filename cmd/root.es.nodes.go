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

package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/trickyearlobe/a2tool/types"
)

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Show data about nodes in the ES cluster",
	Long:  `This command is helpful to understand whether Automate is correctly connected to an external multinode ES cluster`,

	Run: func(cmd *cobra.Command, args []string) {

		// Call the ES REST API
		esURL := fmt.Sprintf("%s://%s:%s/_cat/nodes?pretty&format=json", esScheme, esHost, esPort)
		resp, err := http.Get(esURL)
		if err != nil {
			errorExit("es nodes", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		// Unmarshall the nodes
		var nodes types.ESNodes
		nodes.Unmarshal(body)
		nodes.PrintTable()
	},
}

func init() {
	esCmd.AddCommand(nodesCmd)
}
