package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Abiji-2020/bGuard/config"
//	"github.com/Abiji-2020/bGuard/log"
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
			return newServeCommand().RunE(cmd, args)
		}, SilenceUsage: true,
	}
	c.PersistentFlags().StringVarP(&configPath, "config", "c", defaultConfigPath, "Path to the configuration file")
	c.PersistentFlags().StringVar(&apiHost, "host", defaultHost, "Host to bind the API server")
	c.PersistentFlags().Uint16Var(&apiPort, "port", defaultPort, "Port to bind the API server")

	c.AddCommand(newRefreshCommand(),
		newQueryCommand(),
		newVersionCommand(),
		newServeCommand(),
		newBlockingCommand(),
		newListsCommand(),
		newHealthcheckCommand(),
		newCacheCommand(),
		NewValidateCommand())
	return c

}

func apiURL() string {
	return fmt.Sprintf("http://%s%s", net.JoinHostPort(apiHost, strconv.Itoa(int(apiPort))), "/api")
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

type codeWithStatus interface {
	StatusCode() int
	Status() string
}

func printOkOError(resp codeWithStatus, body string) error {
	if resp.StatusCode() == http.StatusOK {
		log.Log().Info("OK")
	} else {
		return fmt.Errorf("Response NOK, %s %s", resp.Status(), body)

	}
	return nil
}
