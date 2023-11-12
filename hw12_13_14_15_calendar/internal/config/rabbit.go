package config

type RabbitConfig struct {
	User                 string `mapstructure:"RABBIT_USER"`
	Password             string `mapstructure:"RABBIT_PASSWORD"`
	Host                 string `mapstructure:"RABBIT_HOST"`
	Port                 int    `mapstructure:"RABBIT_PORT"`
	Queue                string `mapstructure:"RABBIT_QUEUE"`
	ProducerScanInterval int    `mapstructure:"RABBIT_PRODUCER_SCAN_INTERVAL"`
	CleanerScanInterval  int    `mapstructure:"RABBIT_CLEANER_SCAN_INTERVAL"`
}
