/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"github.com/gedw99/sc-make/assets"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

// listCmd represents the add command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "outputs what files are embedded",
	Long:  `Lists all the embedded files. This is for discovering what templates and examples exist.`,
	Run: func(cmd *cobra.Command, args []string) {

		files, _ := assets.GetAllFilenames()
		for i, s := range files {
			fmt.Println(i, s)
		}

	},
}
