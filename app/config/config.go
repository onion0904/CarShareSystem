package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server Server
	DB     DBConfig
	Mailgun Mailgun
	JWT     JWT
}

type DBConfig struct {
	Name     string `envconfig:"DB_DATABASE" default:"CarShareSystemDB"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASS" default:""`
	Port     string `envconfig:"DB_PORT" default:"3306"`
	Host     string `envconfig:"DB_HOST" default:"db"`
}

type Server struct {
	Address string `envconfig:"ADDRESS" default:"0.0.0.0"`
	Port    string `envconfig:"PORT" default:"8080"`
}

type Mailgun struct {
	Domain string `envconfig:"MAILGUN_DOMAIN"`
	Private_Key string `envconfig:"MAILGUN_PRIVATE_API_KEY"`
	Sender_email string `envconfig:"SENDER_EMAIL"`
	Recipient_email string `envconfig:"RECIPIENT_EMAIL"`
}

type JWT struct {
	Secret string `envconfig:"JWT_SECRET"`
}

var (
	once   sync.Once
	config Config
)

func GetConfig() *Config {
	// goroutine実行中でも一度だけ実行される
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			panic(err)
		}
	})
	return &config
}