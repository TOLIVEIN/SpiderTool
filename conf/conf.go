package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	RedisConf RedisConf `yaml:"redis"`
}

type RedisConf struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Key string `yaml:"key"`
	MaxScore float64 `yaml:"maxScore"`
	MinScore float64 `yaml:"minScore"`
	InitialScore float64 `yaml:"initialScore"`
}

func (c *Config) GetConfig() *Config {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
