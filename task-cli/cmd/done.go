package cmd

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark a task as done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("tasks.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		tasks := strings.Split(string(data), "\n")
		taskNum, err := strconv.Atoi(args[0])
		if err != nil || taskNum <= 0 || taskNum > len(tasks) {
			fmt.Println("Invalid task number")
			return
		}
		tasks[taskNum-1] += " (done)"
		ioutil.WriteFile("tasks.txt", []byte(strings.Join(tasks, "\n")), 0644)
		fmt.Printf("Marked task %d as done\n", taskNum)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
