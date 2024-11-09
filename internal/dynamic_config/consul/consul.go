package consul

import (
	"github.com/ShaleApps/go-service-utils/helpers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	localConsulURL = "localhost:8500"

	AgnusDBConnString        = "agnus_db.conn_string"
	AgnusDBMaxOpenConnection = "agnus_db.max_open_conns"
	AgnusDBMaxIdleConnection = "agnus_db.max_idle_conns"
)

type Consul struct {
	viper *viper.Viper
}

func NewConsulDynamicConfig() Consul {

	viperConfig := viper.New()

	consulURL := helpers.GetEnv("CONSUL_URL", localConsulURL)

	if err := viper.AddRemoteProvider("consul", consulURL, "{{SERVICE_NAME_CAMEL_CASE}}"); err != nil {
		logrus.WithError(err).Error("failed to add remote consul provider for {{SERVICE_NAME_CAMEL_CASE}} viper")
		panic(err)
	}

	viperConfig.SetConfigType("json")

	if err := viperConfig.ReadRemoteConfig(); err != nil {
		logrus.WithError(err).Error("failed to read remote config for {{SERVICE_NAME_CAMEL_CASE}} viper")
	}

	if err := viperConfig.WatchRemoteConfigOnChannel(); err != nil {
		logrus.WithError(err).Error("failed to watch remote config for {{SERVICE_NAME_CAMEL_CASE}} viper")
	}

	return Consul{
		viper: viperConfig,
	}
}

func (cdc Consul) GetAgnusDBConnString() string {
	return cdc.viper.GetString(AgnusDBConnString)
}

func (cdc Consul) GetAgnusDBMaxOpenConnection() int {
	return cdc.viper.GetInt(AgnusDBMaxOpenConnection)
}

func (cdc Consul) GetAgnusDBMaxIdleConnection() int {
	return cdc.viper.GetInt(AgnusDBMaxIdleConnection)
}
