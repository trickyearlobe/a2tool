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
	"github.com/spf13/cobra"
)

// esCmd represents the es command
var esCmd = &cobra.Command{
	Use:   "es",
	Short: "Commands related to Chef Automate ElasticSearch",
}

var esHost string
var esPort string
var esScheme string

func init() {
	rootCmd.AddCommand(esCmd)
	esCmd.PersistentFlags().StringVar(&esHost, "es-host", "localhost", "Hostname of ES load balancer or cluster member")
	esCmd.PersistentFlags().StringVar(&esPort, "es-port", "10141", "Port number of ES load balancer or cluster member")
	esCmd.PersistentFlags().StringVar(&esScheme, "es-scheme", "http", "Scheme can be http or https")
}
