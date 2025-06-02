/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// testeCmd represents the teste command
var testeCmd = &cobra.Command{
	Use:   "teste",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		comando, _ := cmd.Flags().GetString("comando")
		if comando == "ping" {
			cmd.Println("Ping!")
		} else {
			cmd.Println("Pong!")
		}
	},
}

func init() {
	rootCmd.AddCommand(testeCmd)
	testeCmd.Flags().StringP("comando", "c", "", "Escolha ping ou pong")
	testeCmd.MarkFlagRequired("comando")
}
