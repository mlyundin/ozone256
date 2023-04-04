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

type database struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	DBName string `yaml:"db_name"`
}

func (db database) Connection() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Pass, db.DBName)
}

func (s service) Url() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type ConfigStruct struct {
	Services struct {
		Logging struct {
			Devel bool `yaml:"devel"`
		} `yaml:"logging"`

		Checkout struct {
			service     `yaml:",inline"`
			MetricsPort string `yaml:"metrics_port"`
		} `yaml:"checkout"`

		Loms struct {
			service     `yaml:",inline"`
			MetricsPort string `yaml:"metrics_port"`
		} `yaml:"loms"`

		Products struct {
			service `yaml:",inline"`
			Token   string `yaml:"token"`
		} `yaml:"products"`
	} `yaml:"services"`

	Databases struct {
		Checkout database `yaml:"checkout"`
		Loms     database `yaml:"loms"`
	} `yaml:"databases"`

	Kafka struct {
		OrderStatusTopic string    `yaml:"order_status_topic"`
		Brokers          []service `yaml:"brokers,flow"`
	} `yaml:"kafka"`
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
