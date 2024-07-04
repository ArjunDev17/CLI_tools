package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		if _, err := file.WriteString(task + "\n"); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Added task: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
