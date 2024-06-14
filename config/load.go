package config

import "github.com/BurntSushi/toml"

var cfg *Config = DefaultConfig()

func ReadConfig() *Config {
	return cfg
}

func LoadFromTomlFile() error {
	filePath := "D:/IT/Code/goCode/xBlog/config/cfg.toml"
	_, err := toml.DecodeFile(filePath, cfg)
	return err
}
