package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the versions of Go, Java, MongoDB, and MySQL",
	Run: func(cmd *cobra.Command, args []string) {
		checkCommand := func(name string, command []string) {
			cmd := exec.Command(command[0], command[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s version check failed: %v\n", name, err)
				return
			}
			fmt.Printf("%s version:\n%s\n", name, strings.TrimSpace(string(output)))
		}

		checkCommand("Go", []string{"go", "version"})
		fmt.Println("------------------------------------")
		checkCommand("Java", []string{"java", "-version"})
		fmt.Println("------------------------------------")
		checkCommand("MongoDB", []string{"mongod", "--version"})
		fmt.Println("------------------------------------")
		checkCommand("MySQL", []string{"mysql", "--version"})
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
