package cmd

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/database"
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

	api.MustServeFromViper(routers.BridgeRoutes)

}

func init() {
	rootCmd.AddCommand(serverCmd)
}
