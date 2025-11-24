/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		toDos, err := LoadfromJson()
		if err != nil {
			toDos = tasks{}
		}
		toDos.list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
