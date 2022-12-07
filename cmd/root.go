package cmd

import (
	"github.com/create-go-app/cli/v3/pkg/cgapp"
	"github.com/spf13/cobra"
)

var ()

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "gio-make",
	Version: "v",
	Short:   "A powerful CLI for the GIOUI project",
	Long: `
A powerful CLI for the GIOUI project.

Create a new production-ready project with backend (Golang), 
frontend (Golang) and deploy automation 
(Ansible, Docker) by running one CLI command.

-> Focus on writing code and thinking of business logic!
<- The GIO-make CLI will take care of the rest.

A helpful documentation and next steps -> https://github.com/gedw99/gio-make/`,
}

func init() {
	rootCmd.SetOut(cgapp.Stdout)
	rootCmd.SetErr(cgapp.Stderr)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}
