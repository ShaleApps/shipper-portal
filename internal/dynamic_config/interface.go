package dynamic_config

type DynamicConfig interface {
	GetOAuthIssuerURL() string
	GetOAuthAudience() []string
	GetAgnusDBConnString() string
	GetAgnusDBMaxOpenConnection() int
	GetAgnusDBMaxIdleConnection() int
}
