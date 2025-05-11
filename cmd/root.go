package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubetools",
	Short: "A collection of tools for Kubernetes",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(shellCommand)
	rootCmd.AddCommand(logCommand)
}

func initConfig() {
}
