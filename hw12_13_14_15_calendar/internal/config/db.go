package config

type DBConfig struct {
	Type     string `mapstructure:"STORAGE_TYPE"`
	Driver   string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_DB"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
}
