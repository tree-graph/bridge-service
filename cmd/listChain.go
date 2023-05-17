package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
)

// listChainCmd represents the listChain command
var listChainCmd = &cobra.Command{
	Use:   "listChain",
	Short: "list all chains in DB",
	Run: func(cmd *cobra.Command, args []string) {
		database.Init()
		var chains []models.Chain
		if err := database.DB.Find(&chains).Error; err != nil {
			logrus.WithError(err).Error("list chain fail")
		}
		for _, v := range chains {
			logrus.WithField("chain", fmt.Sprintf("%+v", v)).Info("list chain")
		}
	},
}

func init() {
	rootCmd.AddCommand(listChainCmd)
}
