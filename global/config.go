package global

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type DatabaseConfig struct {
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbName"`
	SslMode  string `yaml:"sslMode"`
}

type JwtConfig struct {
	SecretKey       string        `yaml:"secretKey"`
	ExpiresDuration time.Duration `yaml:"expiresDuration"`
}
type Config struct {
	Db  DatabaseConfig `json:"db"`
	Jwt JwtConfig      `json:"jwt"`
}

var (
	config     Config
	configOnce sync.Once
)

func GetConfig() *Config {
	configOnce.Do(func() {
		log.Println("加载配置文件 config.yaml")
		bs, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Fatalln(err)
		}
		_ = yaml.Unmarshal(bs, &config)
	})

	return &config
}
