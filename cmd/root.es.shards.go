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

	"github.com/spf13/cobra"
	"github.com/trickyearlobe/a2tool/types"
)

// shardsCmd represents the shards command
var shardsCmd = &cobra.Command{
	Use:   "shards",
	Short: "Show data about ES shards",
	Run: func(cmd *cobra.Command, args []string) {

		// Call the ES REST API
		esURL := fmt.Sprintf("%s://%s:%s/_cat/shards?format=json", esScheme, esHost, esPort)
		resp, err := http.Get(esURL)
		errorExit(err)

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		// Unmarshall the nodes
		var shards types.ESShards
		shards.Unmarshal(body)
		shards.PrintTable()
	},
}

func init() {
	esCmd.AddCommand(shardsCmd)
}
