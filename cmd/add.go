/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		task, _ := cmd.Flags().GetString("task")
		toDos, err := LoadfromJson()
		if err != nil {
			toDos = tasks{}
		}
		toDos.add(task)
		SavetoJson(toDos)
	},
}

func init() {
	addCmd.Flags().StringP("task", "t", "", "Task to be added")
	rootCmd.AddCommand(addCmd)
}
