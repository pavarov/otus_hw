package config

type AppConfig struct {
	LoggerConfig LoggerConfig     `mapstructure:",squash"`
	DBConfig     DBConfig         `mapstructure:",squash"`
	ServerConfig ServerConfig     `mapstructure:",squash"`
	GrpcConfig   GrpcServerConfig `mapstructure:",squash"`
	RabbitConfig RabbitConfig     `mapstructure:",squash"`
}

func NewAppConfig() AppConfig {
	return AppConfig{}
}
