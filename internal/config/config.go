package config

import (
	"os"
	"path/filepath"

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
	// 尝试从多个位置加载配置文件
	var file []byte
	var err error

	// 1. 当前工作目录
	file, err = os.ReadFile("application.yml")
	if err == nil {
		return parseConfig(file)
	}

	// 2. 程序所在目录
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		configPath := filepath.Join(exeDir, "application.yml")
		file, err = os.ReadFile(configPath)
		if err == nil {
			return parseConfig(file)
		}
	}

	// 3. 父目录
	configPath := filepath.Join("..", "application.yml")
	file, err = os.ReadFile(configPath)
	if err == nil {
		return parseConfig(file)
	}

	return nil, err
}

func parseConfig(file []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(file, &config)
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
