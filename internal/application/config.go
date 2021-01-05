package application

// Config ...
type Config struct {
	RabbitAddr          string `toml:"rabbit_addr"`
	MessageExchangeName string `toml:"rabbit_message_exchange_name"`
	BindAddr            string `toml:"bind_addr"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
