package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Abiji-2020/bGuard/config"
	"github.com/Abiji-2020/bGuard/evt"
	"github.com/Abiji-2020/bGuard/log"
	"github.com/Abiji-2020/bGuard/server"
	"github.com/Abiji-2020/bGuard/util"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var (
	done              = make(chan bool, 1)
	isConfigMandatory = true
	signals           = make(chan os.Signal, 1)
)

func newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "serve",
		Args:              cobra.NoArgs,
		Short:             "start bGuard DNS server (default command)",
		RunE:              startServer,
		PersistentPreRunE: initConfigPreRun,
		SilenceUsage:      true,
	}
}

func startServer(_ *cobra.Command, _ []string) error {
	printBanner()

	cfg, err := config.LoadConfig(configPath, isConfigMandatory)
	if err != nil {
		return fmt.Errorf("unable to load configuration: %w", err)
	}

	log.Configure(&cfg.Log)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	srv, err := server.NewServer(ctx, cfg)
	if err != nil {
		return fmt.Errorf("can't start server: %w", err)
	}

	const errChanSize = 10
	errChan := make(chan error, errChanSize)

	srv.Start(ctx, errChan)

	var terminationErr error

	go func() {
		select {
		case <-signals:
			log.Log().Infof("Terminating...")
			util.LogOnError(ctx, "can't stop server: ", srv.Stop(ctx))
			done <- true

		case err := <-errChan:
			log.Log().Error("server start failed: ", err)
			terminationErr = err
			done <- true
		}
	}()

	evt.Bus().Publish(evt.ApplicationStarted, util.Version, util.BuildTime)
	<-done

	return terminationErr
}


func printBanner() {
	log.Log().Info("_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/")
	log.Log().Info("_/                                                              _/")
	log.Log().Info("_/                                                              _/")
	log.Log().Info("_/       _/                    _/_/_/                           _/")
	log.Log().Info("_/      _/_/_/              _/                                  _/")
	log.Log().Info("_/     _/    _/           _/    _/_/                            _/")
	log.Log().Info("_/    _/    _/            _/    _/                              _/")
	log.Log().Info("_/   _/_/_/                _/_/_/                               _/")
	log.Log().Info("_/                                                              _/")
	log.Log().Info("_/                                                              _/")
	log.Log().Info("_/                                                              _/")
	log.Log().Info("_/_______________________________________________________________/")
	log.Log().Infof("_/  Version: %-18s Build time: %-18s  _/", util.Version, util.BuildTime)
	log.Log().Info("_/                                                              _/")
	log.Log().Info("_/_______________________________________________________________/")

}
