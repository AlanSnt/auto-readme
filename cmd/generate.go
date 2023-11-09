/*
Copyright © 2023 Alan Santer <alansanter46@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const TEMPLATES_FOLDER = "templates"

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new README",
	Long: `Generate a new README with customizable template, based on the project's current state.

You can change the template with the --template flag.

For more information, please visit https://github.com/AlanSnt/auto-readme
	`,
	Run: func(cmd *cobra.Command, args []string) {
		templateSelected, _ := cmd.Flags().GetString("template")
		templateFile := fmt.Sprintf("%s/%s.md", TEMPLATES_FOLDER, templateSelected)

		if templateSelected == "" {
			fmt.Println("No template selected")
			return
		}

		template, err := os.ReadFile(templateFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(template))
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("template", "t", "default", "The template to use")
}
