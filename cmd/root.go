package cmd

import (
	"github.com/Conflux-Chain/go-conflux-util/config"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "bridge",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	blockchain.SetupEvmEnv()
	config.MustInit("")
}
