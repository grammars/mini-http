package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int    `yaml:"port"`
	Log  Log    `yaml:"log"`
	Vhost []Vhost `yaml:"vhost"`
}

type Log struct {
	Folder  string `yaml:"folder"`
	Pattern string `yaml:"pattern"`
}

type Vhost struct {
	Path   string `yaml:"path"`
	Folder string `yaml:"folder"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("application.yml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	if config.Port == 0 {
		config.Port = 80
	}

	if config.Log.Folder == "" {
		config.Log.Folder = "./logs/"
	}

	if config.Log.Pattern == "" {
		config.Log.Pattern = "客户端${IP}通过${METHOD}方法访问了${URL}"
	}

	return &config, nil
}
