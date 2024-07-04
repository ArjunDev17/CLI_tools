package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "version-checker",
	Short: "Version Checker is a CLI tool for checking versions of Go, Java, MongoDB, and MySQL",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
