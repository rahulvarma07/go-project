package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Port string `yaml:"port" env_required:"true"`
}

type Config struct {
	Env        string `yaml:"env" env_required:"true"`
	Storage    string `yaml:"storage" env_required:"true"`
	HttpServer `yaml:"http-server" env_required:"true"`
}

func MustLoadConfig() *Config {
	var getConfigPath string = ""
	// try getting the config path from env

	getConfigPath = os.Getenv("GET_CONFIG_PATH") // getting path from env

	// if getting config path from is not possible
	if getConfigPath == "" {
		// when user gives the argumnest in the CLI take them
		// flag is used to get them

		flags := flag.String("Config", "", "cli given config path")
		flag.Parse()
		getConfigPath = *flags
		// returns and address to dereference we use *

		// make shore to use log if still we dont get any config path
		if getConfigPath == "" {
			log.Fatal("config path is not set")
		}
	}

	// check if the file exsits
	// to return information about file we use Stat
	_, err := os.Stat(getConfigPath)
	if os.IsNotExist(err) {
		log.Fatalf("There is an error loading files in config path %s", getConfigPath) // showing file path if any error
	}

	// if it doe's return the config address
	var finalConfig Config

	// to get values loaded
	// ReadConfig expects two values path and strut(interface)
	err = cleanenv.ReadConfig(getConfigPath, &finalConfig)

	if err != nil {
		log.Fatalf("Failed to read config path %s", err)
	}

	return &finalConfig
}
