package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/helpers"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/infra/evm"
	"github.com/tree-graph/bridge-service/models"
	"strconv"
)

// addReqCmd represents the addReq command
var addReqCmd = &cobra.Command{
	Use:     "addReq",
	Short:   "add a crossing chain request",
	Long:    "args: chainID txHash",
	Args:    cobra.ExactArgs(2),
	Example: "addReq 1 0x0000",
	Run: func(cmd *cobra.Command, args []string) {
		chainId, _ := strconv.ParseInt(args[0], 10, 64)

		database.Init()
		chain, err := models.GetChain(chainId)
		helpers.CheckError("get chain fail:", err)

		evm.AddChainClient(chain)

		_, e := evm.AddRequest(chainId, args[1])
		helpers.CheckError("add request fail:", e)
	},
}

func init() {
	rootCmd.AddCommand(addReqCmd)
}
