/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as completed",
	Long: `Mark a task as completed. For example:
	todo done [id]`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, id := range args {
			i, err := strconv.Atoi(id)
			if err != nil {
				panic("Invalid ID: " + id + ", " + err.Error())
			}
			todoManager.MarkCompleted(i)
			fmt.Printf("Task '%s' marked as completed!\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
