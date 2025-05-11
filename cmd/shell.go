package cmd

import (
	"github.com/spf13/cobra"
)

var shellCommand = &cobra.Command{
	Short: "Opens a pod shell",
	Use:   "shell",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}
