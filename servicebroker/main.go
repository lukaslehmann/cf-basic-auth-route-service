package main

import (
	"os"

	"github.com/pivotal-golang/lager"
)

func main() {
	logger := lager.NewLogger("p-basic-auth-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.ERROR))

	brokerConfigPath := configPath()
}

configPath() string {
	brokerConfigYamlPath := os.Getenv("BROKER_CONFIG_PATH")
	if brokerConfigYamlPath == "" {
		panic("BROKER_CONFIG_PATH not set")
	}
	return brokerConfigYamlPath
}
