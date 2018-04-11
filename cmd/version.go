package cmd

import (
	"fmt"

	"github.com/jdel/gosspks/cfg"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Aliases: []string{
		"v",
	},
	Short: "Get the version of sspks",
	Long:  `Get the version of sspks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gosspks version: ", cfg.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}