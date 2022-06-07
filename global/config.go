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
	DbName   string `yaml:"db-name"`
	SslMode  string `yaml:"ssl-mode"`
}

type JwtConfig struct {
	SecretKey       string        `yaml:"secret-key"`
	ExpiresDuration time.Duration `yaml:"expires-duration"`
}
type ServerConfig struct {
	ListenIp   string `yaml:"listen-ip"`
	ListenPort int    `yaml:"listen-port"`
}
type Config struct {
	Db     DatabaseConfig `yaml:"db"`
	Jwt    JwtConfig      `yaml:"jwt"`
	Server ServerConfig   `yaml:"server"`
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
