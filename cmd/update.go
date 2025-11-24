/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task in the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		index, _ := cmd.Flags().GetInt("index")
		task, _ := cmd.Flags().GetString("task")
		toDos, err := LoadfromJson()
		if err != nil {
			toDos = tasks{}
		}
		toDos.update(index-1, task)
		SavetoJson(toDos)
	},
}

func init() {
	updateCmd.Flags().IntP("index", "i", 0, "Index of the task to update")
	updateCmd.Flags().StringP("task", "t", "", "Task to be updated")
	rootCmd.AddCommand(updateCmd)
}
