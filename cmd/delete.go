/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Task",
	Long: `Delete a task from the list of tasks. For example:
	todo delete [id]`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, id := range args {
			i, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid ID: " + id)
				return
			}
			todoManager.Delete(i)
			fmt.Printf("Task '%s' deleted successfully!\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
