package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	semver  = "0.0.1"
	githash = "HEAD"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display tankerctl version",
	Run:   version,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s (%s)\n", semver, githash)
}
