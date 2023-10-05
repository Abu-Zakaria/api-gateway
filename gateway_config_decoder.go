package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type endpoints struct {
	GET    []string `yaml:"get"`
	POST   []string `yaml:"post"`
	PUT    []string `yaml:"put"`
	DELETE []string `yaml:"delete"`
	PATCH  []string `yaml:"patch"`
}

type service struct {
	Name      string    `yaml:"name"`
	Port      int       `yaml:"port"`
	Base_url  string    `yaml:"base_url"`
	Endpoints endpoints `yaml:"endpoints"`
	SecretKey string    `yaml:"secretKey"`
}

type services struct {
	Services []service `yaml:"services"`
}

func DecodeGatewayConfig() services {
	config_data, err := ioutil.ReadFile("gateway_config.yml")
	if err != nil {
		panic("Something went wrong while trying to read gateway_config.yml file")
	}

	var services services

	err = yaml.Unmarshal(config_data, &services)
	if err != nil {
		panic("Something went wrong while trying to unmarshal gateway_config.yml file")
	}

	return services
}
