package cmd


import (
	"github.com/spf13/cobra"
	"github.com/Reggles44/kube-tools/util/kubernetes"
)

var logCommand = &cobra.Command{
	Use:   "logs",
	Short: "Opens logs",
	Run: func(cmd *cobra.Command, args []string){
		var client = kubernetes.Connect()


	},
}


