package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

// env-default:"production"

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {

	var configPath string

	//Config Path for env
	configPath = os.Getenv("CONFIG_PATH")

	//Config Path from Flags
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags //defrence to the flag variable
	}
	// If still no config path- then error
	if configPath == "" {
		log.Fatal("no config path provided")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist: %s", configPath)
	}

	var cfg Config
	cleanenv.ReadConfig(configPath, &cfg)

}

// env: "dev"
// storage_path : "storage/storage.db"
// http_server:
//   address: "localhost:8082"
