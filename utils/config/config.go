package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// AppConfig 应用基础配置
type AppConfig struct {
	AppName   string `json:"name" yaml:"name"`
	AppEnv    string `json:"env" yaml:"env"`
	AppPort   int    `json:"port" yaml:"port"`
	DebugPort int    `json:"debug_port" yaml:"debug_port"`
}

type Database struct {
	Driver   string `json:"driver"`
	Port     int    `json:"port"`
	Address  string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Auth     string `json:"auth"`
	PoolSize int    `json:"pool_size"`
}

type Configuration struct {
	App   AppConfig
	Db    Database
	Redis Redis
}

var once sync.Once
var config *Configuration

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("/Users/tiejiaxiaobao/GolandProjects/OJ/config/config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return config
}
