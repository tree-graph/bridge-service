package cmd

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/infra/worker"
	"github.com/tree-graph/bridge-service/routers"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run web server",
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func main() {
	database.Init()
	if err := worker.SetupChains(); err != nil {
		logrus.WithError(err).Fatal("setup chains fail")
	}
	api.MustServeFromViper(routers.BridgeRoutes)

}

func init() {
	rootCmd.AddCommand(serverCmd)
}
