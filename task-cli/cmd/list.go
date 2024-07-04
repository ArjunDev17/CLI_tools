package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("tasks.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		tasks := strings.Split(string(data), "\n")
		fmt.Println("Tasks:")
		for i, task := range tasks {
			if task != "" {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
