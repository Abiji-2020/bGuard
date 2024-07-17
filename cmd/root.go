package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	configPath string
	apiHost    string
	apiPort    uint16
)

const (
	defaultPort        = 4000
	defaultHost        = "localhost"
	defaultConfigPath  = "./config.yaml"
	configFileEnVar    = "BGUARD_CONFIG_FILE"
	configFileEnVarOld = "BGUARD_CONFIG"
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

func initConfig() error {
	if configPath == defaultConfigPath {
		val, present := os.LookupEnv(configFileEnVar)
		if present {
			configPath = val
		} else {
			val, present = os.LookupEnv(configFileEnVarOld)
			if present {
				configPath = val
			}
		}
	}

	cfg, err := config.LoadConfig(configPath, false)
	if err != nil {
		return fmt.Errorf("Unable to load configuration file %s: %w", configPath, err)
	}
	log.confiure(&cfg.Log)
	if len(cfg.Ports.HTTP) != 0 {
		split := strings.Split(cfg.Ports.HTTP, ":")
		lastIdx := len(split) - 1
		apiHost = strings.Join(split[:lastIdx], ":")
		port, err := config.ConvertPort(split[lastIdx])
		if err != nil {
			return fmt.Errorf("Unable to parse port number %s: %w", split[lastIdx], err)
		}
		apiPort = apiPort
	}
	return nil

}
