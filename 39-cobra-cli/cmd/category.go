/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := cmd.Flags().GetString("name")
		fmt.Println("Category command executed with name:" + category)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("Category exists:", exists)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PreRun: This will run before the main command logic.")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PostRun: This will run after the main command logic.")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("an error occurred while executing the command")
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.Flags().StringVarP(&category, "name", "n", "", "Name of the category")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if the category exists")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
