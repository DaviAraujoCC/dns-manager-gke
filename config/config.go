package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	ProjectId          string `mapstructure:"project_id"`
	ManagedZone        string `mapstructure:"managed_zone"`
	DnsSuffix          string `mapstructure:"dns_suffix"`
	Namespace          string `mapstructure:"namespace"`
	IgnoreDeleteRecord bool   `mapstructure:"ignore_delete_record"`
}

func New() (*Config, error) {
	viper.BindEnv("PROJECT_ID")
	viper.BindEnv("MANAGED_ZONE")
	viper.BindEnv("DNS_SUFFIX")
	viper.SetDefault("NAMESPACE", "default")
	viper.SetDefault("IGNORE_DELETE_RECORD", false)
	viper.AutomaticEnv()
	cfg := &Config{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.DnsSuffix == "" {
		return nil, errors.New("DNS_SUFFIX environment variable is not set")
	}

	if cfg.ProjectId == "" {
		return nil, errors.New("PROJECT_ID environment variable is not set")
	}

	if cfg.ManagedZone == "" {
		return nil, errors.New("MANAGED_ZONE environment variable is not set")
	}

	return cfg, nil
}
