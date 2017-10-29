package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is the root command of tankerctl
var RootCmd = &cobra.Command{
	Use:   "tankerctl",
	Short: "Export gasoline data as sensision metrics",
}

func init() {
	cobra.OnInitialize(configure)

	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")

	viper.BindPFlags(RootCmd.PersistentFlags())
}

func configure() {
	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	viper.AddConfigPath("/etc/tankerctl")
	viper.AddConfigPath("$HOME/.tankerctl")
	viper.AddConfigPath(".")

	viper.SetConfigName("config")

	if err := viper.MergeInConfig(); err != nil {
		log.Warn(err)
	}
}
