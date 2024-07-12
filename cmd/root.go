package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	c := &cobra.Command{
		Use:    "bGuard",
		Short:  "bGuard is a simple CLI tool for managing your DNS",
		Long:   "bGuard is a simple CLI tool for managing your DNS by providing custom DNS resolver and ad-blocker to restrict the junk and unwanted traffic",
		PreRun: initConfigPreRun,
		RunE: func(cmd *cobra.Command, args []string) error {
			return newServerCommand().RunE(cmd, args)
		}, SilenceUsage: true,
	}

}

func Execute() {
	if err := NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfigPreRun(cmd *cobra.Command, args []string) {
	return initConfig()
}
