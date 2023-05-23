package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/infra/worker"
	"os"
	"os/signal"
	"syscall"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "cross-chain event worker",
	Run: func(cmd *cobra.Command, args []string) {
		database.Init()
		runClaim, err := cmd.Flags().GetBool("claim")
		if err != nil {
			logrus.WithError(err).Fatal("parse flag <claim> fail")
			return
		}
		if runClaim {
			if err := worker.RunAllClaimWorker(); err != nil {
				logrus.WithError(err).Error("RunAllClaimWorker fail")
				return
			}
		} else if err := worker.RunAllCrossEventWorker(); err != nil {
			logrus.WithError(err).Error("RunAllCrossEventWorker fail")
			return
		}
		blockMainThread()
	},
}

func blockMainThread() {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.Flags().BoolP("claim", "c", false, "run claim worker")
}
