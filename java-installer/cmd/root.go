package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "java-installer",
	Short: "A CLI tool to install Java",
	Run: func(cmd *cobra.Command, args []string) {
		commands := []string{
			"sudo apt update",
			"sudo apt install -y openjdk-17-jdk",
			"java -version",
		}

		for _, cmdStr := range commands {
			fmt.Println("Running:", cmdStr)
			cmd := exec.Command("bash", "-c", cmdStr)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running command %s: %v\n", cmdStr, err)
				return
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define flags and configuration settings.
}
