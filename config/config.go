package config

import "github.com/ilyakaznacheev/cleanenv"

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type InfluxInstance struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Token    string `yaml:"token"`
}
type Config struct {
	Server         Server         `yaml:"server"`
	InfluxInstance InfluxInstance `yaml:"influxInstance"`
}

func LoadConfiguration() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
