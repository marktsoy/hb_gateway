package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/marktsoy/hb_gateway/internal/application"
)

var (
	configPath string
)

// init function - called once
// More info: https://tutorialedge.net/golang/the-go-init-function/#:~:text=The%20init%20Function,will%20only%20be%20called%20once.
func init() {
	fmt.Println("Init function. I am called only once. I am called first")

	// Flag (cmd args) definitions
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "Path to config file")
}

func main() {
	// After all flags are defined call the Parse function to bind args to vars
	flag.Parse()
	var config *application.Config = application.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rabbit Addr: %v \n", config.RabbitAddr)
	fmt.Printf("Rabbit Message Queue: %v \n", config.MessageQueueName)

	server := application.New(config)

	server.Run()

}
