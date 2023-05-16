/*

 */
package cmd

import (
	"github.com/Conflux-Chain/go-conflux-util/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "bridge",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	config.MustInit("")
	operatorSysLogDir := viper.GetString("log.file")

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(operatorSysLogDir),
		MaxSize:    5, // MB
		MaxBackups: 10,
		MaxAge:     30, // days
		Compress:   false,
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, lumberjackLogger))
	logrus.Infof("operatorSysLogDir [%s] \n", operatorSysLogDir)
}
