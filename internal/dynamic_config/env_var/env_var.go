package env_var

import (
	"github.com/ShaleApps/go-service-utils/helpers"
)

type EnvVar struct {
}

func NewEnvVarDynamicConfig() *EnvVar {
	return &EnvVar{}
}

func (cdc EnvVar) GetAgnusDBConnString() string {
	return helpers.GetEnv("DB_URL", "host=localhost port=54325 user=postgres dbname={{SERVICE_NAME_SNAKE_CASE}} password=secret sslmode=disable TimeZone=UTC")
}

func (cdc EnvVar) GetAgnusDBMaxOpenConnection() int {
	return helpers.GetIntEnv("DB_MAX_OPEN_CONNECTION", 20)
}

func (cdc EnvVar) GetAgnusDBMaxIdleConnection() int {
	return helpers.GetIntEnv("DB_MAX_IDLE_CONNECTION", 20)
}
