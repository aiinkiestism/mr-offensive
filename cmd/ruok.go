/*
Copyright © 2020 aiinkiestism <hashmimic1@gmail.com>

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

	"github.com/spf13/cobra"
)

// ruokCmd represents the ruok command
var ruokCmd = &cobra.Command{
	Use:   "ruok",
	Short: "respond to ruok",
	Long:  `Long help doesn't help you. You just need to try.`,
	// `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kidding me? You son of a b***h")
	},
}

func init() {
	rootCmd.AddCommand(ruokCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ruokCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ruokCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
