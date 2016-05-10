package config

import (
	"os"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type Config struct {
	BrokerConfiguration BrokerConfiguration `yaml:"basic_auth_service_broker"`
}

type BrokerConfiguration struct {
	RouteServiceURL string `yaml:"route_service_url"`
	BrokerUserName  string `yaml:"broker_username"`
	BrokerPassword  string `yaml:"broker_password"`
}

func ParseConfig(path string) (BrokerConfiguration, error) {
	file, err := os.Open(path)
	if err != nil {
		return BrokerConfiguration{}, err
	}

	var config BrokerConfiguration
	if err := candiedyaml.NewDecoder(file).Decode(&config); err != nil {
		return BrokerConfiguration{}, err
	}

	return config, nil
}
