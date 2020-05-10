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
)

var level string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the ES cluster status at varying levels of detail",
	Run: func(cmd *cobra.Command, args []string) {
		esURL := fmt.Sprintf("%s://%s:%s/_cluster/health?pretty&level=%s", esScheme, esHost, esPort, level)
		resp, err := http.Get(esURL)
		if err != nil {
			errorExit("es status", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func init() {
	esCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringVar(&level, "level", "cluster", "Level of detail can be cluster, indices or shards")
}
