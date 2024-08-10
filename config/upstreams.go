package config

import (
	"github.com/Abiji-2020/bGuard/log"
	"github.com/sirupsen/logrus"
)

const UpstreamDefaultCfgName = "default"

type Upstreams struct {
	Init      Init             `yaml:"init"`
	Timeout   Duration         `yaml:"timeout" default:"2s"`
	Groups    UpstreamGroups   `yaml:"groups"`
	Strategy  UpstreamStrategy `yaml:"strategy" default:"parallel_best"`
	UserAgent string           `yaml:"user_agent"`
}

type UpstreamGroups map[string][]Upstream

func (c *Upstreams) validate(logger *logrus.Entry) {
	defaults := mustDefault[Upstreams]()
	if !c.Timeout.IsAboveZero() {
		logger.Warnf("Upstreams.timeout <= 0, setting to %s", defaults.Timeout)
		c.Timeout = defaults.Timeout
	}
}

func (c *Upstreams) IsEnabled() bool {
	return len(c.Groups) != 0
}

func (c *Upstreams) LogConfig(logger *logrus.Entry) {
	logger.Info("Init:")
	log.WithIndent(logger, " ", c.Init.LogConfig)

	logger.Info("timeout: ", c.Timeout)
	logger.Info("strategy: ", c.Strategy)
	logger.Info("Groups:")
	for name, upstreams := range c.Groups {
		logger.Infof("  %s:", name)
		for _, u := range upstreams {
			logger.Infof("       - %s", u)
		}
	}
}

type UpstreamGroup struct {
	Upstreams
	Name string
}

func NewUpstreamGroup(name string, cfg Upstreams, upstreams []Upstream) UpstreamGroup {
	group := UpstreamGroup{
		Name:      name,
		Upstreams: cfg,
	}
	group.Groups = UpstreamGroups{name: upstreams}
	return group
}

func (c *UpstreamGroup) GroupUpstreams() []Upstream {
	return c.Groups[c.Name]
}

func (c *UpstreamGroup) IsEnabled() bool {
	return len(c.GroupUpstreams()) != 0
}

func (c *UpstreamGroup) LogConfig(logger *logrus.Entry) {
	logger.Info("Group: ", c.Name)
	logger.Info("Upstreams:")
	for _, u := range c.GroupUpstreams() {
		logger.Infof("       - %s", u)
	}
}
