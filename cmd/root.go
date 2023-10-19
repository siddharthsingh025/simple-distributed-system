// cmd/root.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "distributedSystemCLI",
	Short: "A simple CLI-based distributed system",
	Long:  `A simple CLI-based distributed system with a basic UI submodule.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Simple CLI-based Distributed System!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
