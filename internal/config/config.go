package config

import (
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/dynamic_config"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/dynamic_config/consul"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/dynamic_config/env_var"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/metrics/prometheus"
	"github.com/ShaleApps/go-service-utils/helpers"
	"github.com/sirupsen/logrus"
)

type SvcConfig struct {
	MetricsCollector prometheus.Collectors
	DynamicConfig    dynamic_config.DynamicConfig
}

func LoadConfig() SvcConfig {

	cfg := SvcConfig{}

	cfg.MetricsCollector = prometheus.NewMetricsCollector()

	cfg.MetricsCollector.Register()

	initDynamicConfigDriver(&cfg)

	return cfg
}

func initDynamicConfigDriver(cfg *SvcConfig) {

	switch helpers.GetEnv("DYNAMIC_CONFIG_DRIVER", "") {
	case "consul":
		logrus.Info("Using Consul dynamic config driver")
		cfg.DynamicConfig = consul.NewConsulDynamicConfig()
	default:
		logrus.Info("No dynamic config driver defined, using env_var driver")
		cfg.DynamicConfig = env_var.NewEnvVarDynamicConfig()
	}
}
