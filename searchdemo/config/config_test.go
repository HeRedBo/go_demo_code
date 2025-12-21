package config

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadConfig()
	dump.P(Cfg)
	dump.P(Cfg.App)
	dump.P(Cfg.Mysql)
}
