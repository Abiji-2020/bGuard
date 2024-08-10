package config

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/miekg/dns"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	. "github.com/Abiji-2020/bGuard/config/migration"
	"github.com/Abiji-2020/bGuard/log"
	"github.com/Abiji-2020/bGuard/util"
)

const (
	updPort   = 53
	tlsPort   = 853
	httpsPort = 443

	secretObfuscator = "********"
)

type Configurable interface {
	isEnabled() bool
	LogConfig(*logrus.Entry)
}

type NetProtocol uint16

type IPVersion uint8

func (ipv IPVersion) Net() string {
	switch ipv {
	case IPVersionDual:
		return "ip"
	case IPVersionV4:
		return "ip4"
	case IPVersionV6:
		return "ip6"
	}

	panic(fmt.Errorf("bad value: %s", ipv))
}

func (ipv IPVersion) QTypes() []dns.Type {
	switch ipv {
	case IPVersionDual:
		return []dns.Type{dns.TypeA, dns.TypeAAAA}
	case IPVersionV4:
		return []dns.Type{dns.TypeA}
	case IPVersionV6:
		return []dns.Type{dns.TypeAAAA}
	}

	panic(fmt.Errorf("bad value: %s", ipv))
}

type TLSVersion int

func (v *TLSVersion) validate(logger *logrus.Entry) {

	minAllowed := tls.config{MinVersion: tls.VersionTLS12}.MinVersion

	if *v < TLSVersion(minAllowed) {
		def := mustDefault[Config]().minTLSServeVer

		logger.Warnf("TLS version %s is not supported, using %s instead", v, def)
		*v = def
	}
}

type QueryLogType int16

type InitStrategy uint16

func (s InitStrategy) Do(ctx context.Context, init func(context.Context) error, logErr func(error)) error {
	init = recoverToError(init, func(panicVal any) error {
		return fmt.Errorf("panic during initalization: %v", panicVal)
	})

	if s == InitStrategyFast {
		go func() {
			err := init(ctx)
			if err != nil {
				logErr(err)
			}
		}()
		return nil
	}

	err := init(ctx)
	if err != nil {
		logErr(err)

		if s == InitStrategyFailOnError {
			return err
		}

	}
	return nil
}

type QueryLogField string

type UpstreamStrategy uint8

var netDefaultPort = map[NetProtocol]uint16{
	NetProtocolTcpUdp: updPort,
	NetProtocolTcpTls: tlsPort,
	NetProtocolHttps:  httpsPort,
}

type ListenConfig []string

func (l *ListenConfig) UnmarshalText(data []byte) error {
	addresses := string(data)

	*l = strings.Split(addresses, ",")

	return nil
}

func (b *BootstrapDNS) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var single BootStrappedUpstream

	if err := unmarshal(&single); err == nil {
		*b = BootstrapDNS{single}

		return nil
	}

	var c BootstrapDNS
	if err := unmarshal(&c); err != nil {
		return err
	}

	*b = BootstrapDNS(c)
	return nil
}


func (b *BootStrappedUpstream) UnmarshalYAML (unmarshal func (interface{}) error) error {
	if err := unmarshal(&b.Upstream); err == nil {
		return nil
	}

	// bootstrappedUpstream is used to avoid infinite recursion:
	// if we used BootstrappedUpstream, unmarshal would just call us again.
	var c bootstrappedUpstream
	if err := unmarshal(&c); err != nil {
		return err
	}

	*b = BootstrappedUpstream(c)

	return nil
}


type Config struct {
	Upstreams        Upstreams           `yaml:"upstreams"`
	ConnectIPVersion IPVersion           `yaml:"connectIPVersion"`
	CustomDNS        CustomDNS           `yaml:"customDNS"`
	Conditional      ConditionalUpstream `yaml:"conditional"`
	Blocking         Blocking            `yaml:"blocking"`
	ClientLookup     ClientLookup        `yaml:"clientLookup"`
	Caching          Caching             `yaml:"caching"`
	QueryLog         QueryLog            `yaml:"queryLog"`
	Redis            Redis               `yaml:"redis"`
	Log              log.Config          `yaml:"log"`
	Ports            Ports               `yaml:"ports"`
	MinTLSServeVer   TLSVersion          `yaml:"minTlsServeVersion" default:"1.2"`
	CertFile         string              `yaml:"certFile"`
	KeyFile          string              `yaml:"keyFile"`
	BootstrapDNS     BootstrapDNS        `yaml:"bootstrapDns"`
	HostsFile        HostsFile           `yaml:"hostsFile"`
	FQDNOnly         FQDNOnly            `yaml:"fqdnOnly"`
	Filtering        Filtering           `yaml:"filtering"`
	EDE              EDE                 `yaml:"ede"`
	ECS              ECS                 `yaml:"ecs"`
	SUDN             SUDN                `yaml:"specialUseDomains"`

	// Deprecated options
	Deprecated struct {
		Upstream            *UpstreamGroups `yaml:"upstream"`
		UpstreamTimeout     *Duration       `yaml:"upstreamTimeout"`
		DisableIPv6         *bool           `yaml:"disableIPv6"`
		LogLevel            *logrus.Level   `yaml:"logLevel"`
		LogFormat           *log.FormatType `yaml:"logFormat"`
		LogPrivacy          *bool           `yaml:"logPrivacy"`
		LogTimestamp        *bool           `yaml:"logTimestamp"`
		DNSPorts            *ListenConfig   `yaml:"port"`
		HTTPPorts           *ListenConfig   `yaml:"httpPort"`
		HTTPSPorts          *ListenConfig   `yaml:"httpsPort"`
		TLSPorts            *ListenConfig   `yaml:"tlsPort"`
		StartVerifyUpstream *bool           `yaml:"startVerifyUpstream"`
		DoHUserAgent        *string         `yaml:"dohUserAgent"`
	} `yaml:",inline"`
}

type Ports struct {
	DNS ListenConfig `yaml:"dns" default: "53"`

	HTTP ListenConfig `yaml:"http"` 
	HTTPS ListenConfig `yaml:"https"` 
	TLS ListenConfig `yaml:"tls"`
}

func (c *Ports) LogConfig(logger *logrus.Entry){
	logger.Infof("DNS port: %s", c.DNS)
	logger.Infof("HTTP port: %s", c.HTTP)
	logger.Infof("HTTPS port: %s", c.HTTPS)
	logger.Infof("TLS port: %s", c.TLS)
}

type (
	BootstrapDNS bootstrapDNS 
	bootstrapDNS []BootstrappedUpstream

)

func (b *BootstrapDNS) isEnabled() bool {
	return len(*b) != 0
}

func (b *BootstrapDNS) LogConfig(*logrus.Entry) {
	panic("Not implemented")
}

type (
	BootstrappedUpstream bootstrappedUpstream
	bootstrappedUpstream struct {
		Upstream Upstream `yaml:"upstream"`
		IPs []net.IP `yaml:"ips"`
	}
)

type (
	FQDNOnly = toEnable
	EDE = toEnable
)


type toEnable struct {
	Enable bool `yaml:"enable" default: false`
}


func (e *toEnable) isEnabled() bool {
	return e.Enable
}

func (c *toEnable) LogConfig(logger *logrus.Entry) {
	logger.Infof("Enabled")
}

type Init struct {
	Strategy InitStrategy `yaml:"strategy" default:"blocking"`
}

func (c *Init) LogConfig(logger *logrus.Entry) {
	logger.Debugf("Strategy =  %s", c.Strategy)
}

type SourceLoading struct{
	Init `yaml:",inline"`

	Concurrency uint `yaml:"concurrency" default:"4"`
	MaxErrorsPerSource int `yaml:"maxErrorsPerSource" default:"5"`
	RefreshPeriod Duration `yaml:"refreshPeriod" default:"3h"`
	Downloads Downloader `yaml:"downloads"`
}

func (c *SourceLoading) LogConfig(logger *logrus.Entry){
	c.Init.LogConfig(logger)
	logger.Infof("Concurrency = %d", c.Concurrency)
	logger.Debugf("Max errors per source = %d", c.MaxErrorsPerSource)

	if c.RefreshPeriod.IsAboveZero(){
		logger.Infof("Refresh  = every %s", c.RefreshPeriod)
	} else {
		logger.Infof("Refresh = disabled")
	}

	logger.Info("Downloads:")
	log.WithIndent(logger, " ", c.Downloads.LogConfig)

}

func (c *SourceLoading) StartPeriodicRefresh(ctx context.Context, refresh func(context.Context) error, logErr func(error),) error{
	err := c.Strategy.Do(ctx, refresh, logErr)

	if err != nil {
		return err
	}

	if c.RefreshPeriod > 0 {
		go c.periodically(ctx, refresh, logErr)
	}
	

	return nil
}


func (c *SourceLoading) periodically( ctx context.Context, refresh func(context.Context) error, logErr func(error),){
	refresh = recoverToError(refresh, func(panicVal any) error{
		return fmt.Errorf("panic during periodic refresh: %v", panicVal)
	}

	ticker := time.NewTicker(c.RefreshPeriod.ToDuration())

	defer ticker.Stop()

	for {
		select {
			case <-ticker.C:
				err := refresh(ctx)
				if err != nil {
					logErr(err)
				}
			case <-ctx.Done():
				return

		}
	}
}


func recoverToError(do func(context.Context) error, onPanic func(any) error) func(context.Context) error {
	return func(ctx context.Context) error {
		defer func(){
			if val:= recover(); val != nil {
				rerr = onPanic(val)
			}
		}()

		return do(ctx)
	}
}

type Downloader struct {
	Timeout Duration `yaml:"timeout" default:"5s"`
	Attempts uint `yaml:"attempts" default:"3"`
	Cooldown Duration `yaml:"cooldown" default:"500ms"`
}


func (d *Downloader) LogConfig(logger *logrus.Entry){
	logger.Infof("Timeout = %s", d.Timeout)
	logger.Infof("Attempts = %d", d.Attempts)
	logger.Infof("Cooldown = %s", d.Cooldown)
}

func WithDefaults[T any]()(T, error){
	var cfg T 
	if err:= defaults.Set(&cfg); err!= nil{
		reurn cfg, fmt.Errorf("can't apply %T defaults: %w", cfg, err)
	}

	return cfg, nil
}

func mustDefault[T any]() T {
	cfg, err := WithDefaults[T]()
	if err != nil {
		util.FatalOnError("broken defaults", err)
	}

	return cfg
}

func LoadConfig(path string,mandatory bool)(rCfg *Config, rerr error){
	logger := logrus.NewEntry(log.Log())

	return loadConfig(logger, path, mandatory)
}

