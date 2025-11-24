/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// markCompletedCmd represents the markCompleted command
var markCompletedCmd = &cobra.Command{
	Use:   "markCompleted",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		index, _ := cmd.Flags().GetInt("index")
		toDos, err := LoadfromJson()
		if err != nil {
			toDos = tasks{}
		}
		if index > 0 && index <= len(toDos) {
			toDos.markCompleted(index - 1)
			SavetoJson(toDos)
		}
	},
}

func init() {
	markCompletedCmd.Flags().IntP("index", "i", 0, "Index of the task to mark as completed")
	rootCmd.AddCommand(markCompletedCmd)
}
