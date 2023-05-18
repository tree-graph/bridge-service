package cmd

import (
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/infra/worker"

	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "cross-chain request worker",
	Run: func(cmd *cobra.Command, args []string) {
		database.Init()
		worker.InitCrossReqWorker()
		worker.CrossRequestWorker()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
