package config

type GrpcServerConfig struct {
	Port int `mapstructure:"GRPC_SERVER_PORT"`
}
