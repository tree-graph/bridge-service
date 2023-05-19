package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/infra/worker"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "cross-chain event worker",
	Run: func(cmd *cobra.Command, args []string) {
		database.Init()
		if err := worker.RunAllCrossEventWorker(); err != nil {
			logrus.WithError(err).Error("RunAllCrossEventWorker fail")
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
