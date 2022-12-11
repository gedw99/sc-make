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
	"log"
	"os"

	"github.com/gedw99/sc-make/assets"
	"github.com/spf13/cobra"

	"github.com/FalcoSuessgott/gitget/repo"
)

func init() {
	rootCmd.AddCommand(createCmd)

	// Git is the remote git repo we are looking to use.
	createCmd.Flags().StringP("git", "g", "https://github.com/gioui/gio-example", "Git URL to use")

	// tempate can be the embedded one inside Projects OR a remote one.
	// for now we just look for the embedded one, as thats all we need.
	createCmd.Flags().StringP("template", "t", "https://github.com/spf13/cobra", "Template to use")

	//createCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
	//createCmd.Flags().StringP("gio-template1", "p", "https://github.com/spf13/cobra", "GIO Template to use")
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a local project from a git repository",
	Long: `Create will create a git project locally with the template, so that you can immediately work with the code.

For Example: gio-make create --git https://github.com/gioui/gio-example --template`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println()

		fmt.Println("-- Running Create task --")

		// File path we are on.
		cwdPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("File path: " + cwdPath)
		//pretty.Sprint("File path: ", cwdPath)

		//fstatus, _ := cmd.Flags().GetBool("float")
		git, _ := cmd.Flags().GetString("git")
		fmt.Println("git: " + git)

		template, _ := cmd.Flags().GetString("template")
		fmt.Println("template: " + template)

		// Use current user runtime dir
		targetDir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		// step 1: Extract the example to the FS.
		err = assets.CreateProjectFiles(targetDir, "projname", git, template)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		// step 2: Extract the template ( git is defualt) .
		// first we need to know the git repo being used, then we can find / replace the makefile
		if git != "" {
			fmt.Println("extracting git fragments")
			fmt.Println("git url:  " + git)

			gitIsUrl := repo.IsGitURL(git)
			fmt.Printf("git is valid: %v ", gitIsUrl)

			//repo := repo.NewRepository(git)
			//fmt.Print(repo.Files)

		}

	},
}
