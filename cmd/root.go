/*
Copyright © 2023 Alan Santer <alansanter46@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auto-readme",
	Short: "CLI tool to effortlessly generate project READMEs using customizable templates",
	Long: `ProjectAutoReadme is a robust command-line interface (CLI) built for MAC OS and Linux.

Simplify your project documentation process with customizable templates, seamlessly integrated into your workflow.

ProjectAutoReadme streamlines the creation of essential documentation, saving you time and ensuring consistency.
Enhance your project management with this user-friendly CLI. Get started today!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
