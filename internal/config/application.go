package config

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/pkg/errors"
)

type AppPostgreSQLConfig struct {
	URL      string `hcl:"url,optional"`
	Host     string `hcl:"host,optional"`
	Port     string `hcl:"port,optional"`
	Database string `hcl:"database,optional"`
	Username string `hcl:"username,optional"`
	Password string `hcl:"password,optional"`
}

// AppConfig represents a configuration format for the Response application.
type AppConfig struct {
	EncryptionKey string               `hcl:"encryption_key,attr"`
	PostgreSQL    *AppPostgreSQLConfig `hcl:"postgresql,block"`
}

// ParseApplicationConfig parses an application configuration file and returns the parsed
// result.
func ParseApplicationConfig(config string) (*AppConfig, error) {
	appConfig := &AppConfig{}

	if err := hclsimple.DecodeFile(config, CreateContext(), appConfig); err != nil {
		return nil, errors.Wrap(err, "unable to parse application config")
	}

	return appConfig, nil
}
