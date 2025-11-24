/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// clearListCmd represents the clearList command
var clearListCmd = &cobra.Command{
	Use:   "clearList",
	Short: "Clear the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		toDos, err := LoadfromJson()
		if err != nil {
			toDos = tasks{}
		}
		toDos.clearList()
		SavetoJson(toDos)
	},
}

func init() {
	rootCmd.AddCommand(clearListCmd)
}
