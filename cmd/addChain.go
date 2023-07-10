package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"strconv"
)

// addChainCmd represents the addChain command
var addChainCmd = &cobra.Command{
	Use:   "addChain",
	Short: "Add a chain",
	Long:  `Add a chain, args: chainId name rpc <evm|cfx>`,
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) {
	database.Init()
	chainId, _ := strconv.ParseInt(args[0], 10, 64)
	var bean = models.Chain{
		ChainId: chainId, Name: args[1], Rpc: args[2], ChainType: args[3],
	}
	if err := database.GetDB().Create(&bean).Error; err != nil {
		logrus.WithError(err).Error("config failed")
	}
	logrus.Info("created")
}

func init() {
	rootCmd.AddCommand(addChainCmd)
}
