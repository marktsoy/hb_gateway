package application

// Config ...
type Config struct {
	RabbitAddr       string `toml:"rabbit_addr"`
	MessageQueueName string `toml:"rabbit_message_queue"`
	BindAddr         string `toml:"bind_addr"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
