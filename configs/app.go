package configs

import (
	"sync"
)

type AppConfig struct {
	name    string
	version string
	env     string
	host    string
	schemes []string
	storage string
}

var (
	app     *AppConfig
	appOnce sync.Once
)

func (ac *AppConfig) Name() string {
	return ac.name
}

func (ac *AppConfig) Version() string {
	return ac.version
}

func (ac *AppConfig) Env() string {
	return ac.env
}

func (ac *AppConfig) Host() string {
	return ac.host
}

func (ac *AppConfig) Schemes() []string {
	return ac.schemes
}

func (ac *AppConfig) StoragePath() string {
	return ac.storage
}

func App() *AppConfig {
	appOnce.Do(func() {
		app = &AppConfig{
			name:    "mcash-finance-consolidation-core",
		}
	})
	return app
}
