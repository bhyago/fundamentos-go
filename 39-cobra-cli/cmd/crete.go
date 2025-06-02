/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cobracli/internal/database"

	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  "Create a new category with a name and description.",
		RunE:  runCreate(categoryDb),
	}
	return cmd
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}

		cmd.Println("Category created successfully!")
		return nil
	}
}

func init() {
	creteCmd := newCreateCmd(GetCategoryDB(GetDb()))
	creteCmd.Flags().StringP("name", "n", "", "Name of the category")
	creteCmd.Flags().StringP("description", "d", "", "Description of the category")
	creteCmd.MarkFlagRequired("name")
	creteCmd.MarkFlagRequired("description")
	creteCmd.MarkFlagsRequiredTogether("name", "description")
}
