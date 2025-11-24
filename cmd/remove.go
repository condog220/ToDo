/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task from the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		index, _ := cmd.Flags().GetInt("index")
		toDos, err := LoadfromJson()
		if err != nil {
			toDos = tasks{}
		}
		toDos.remove(index - 1)
		SavetoJson(toDos)
	},
}

func init() {
	removeCmd.Flags().IntP("index", "i", 0, "Index of the task to remove")
	rootCmd.AddCommand(removeCmd)
}
