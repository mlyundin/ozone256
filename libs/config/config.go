package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type service struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s service) Url() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type ConfigStruct struct {
	Services struct {
		Checkout service `yaml:"checkout"`
		Loms     service `yaml:"loms"`
		Products struct {
			Service service `yaml:"service"`
			Token   string  `yaml:"token"`
		} `yaml:"products"`
	} `yaml:"services"`
}

var ConfigData ConfigStruct

func Init() error {
	rawYAML, err := os.ReadFile("config.yml")
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}
