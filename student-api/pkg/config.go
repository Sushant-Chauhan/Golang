package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

// env-default:"production"

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Fatal("config file doesn't exist")
		}

		var cfg Config
		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			log.Fatal("can not read config file: %s", err.Error())
		}

		return &cfg

	}
}
