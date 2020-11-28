package config

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config wraps all necessary parameters
type Config struct {
	Node NodeConfig `yaml:"node"`
	DB   DBConfig   `yaml:"database"`
}

// NodeConfig wraps all node endpoints that are used in this project
type NodeConfig struct {
	ExchangeAPIEndpoint string `yaml:"exchange_api_endpoint"`
	ChainID             string `yaml:"chain_id"`
}

// DBConfig wraps all required parameters for database connection
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Table    string `yaml:"table"`
}

// ParseConfig attempts to read and parse config.yaml from the given path
// An error reading or parsing the config results in a panic.
func ParseConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../") // for test cases

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config"))
	}

	cfg := Config{}

	if viper.GetString("active") == "" {
		log.Fatal("define active chain_id param in your config file.")
	}

	switch viper.GetString("active") {
	case "888":
		cfg.Node = NodeConfig{
			ExchangeAPIEndpoint: viper.GetString("testnet.node.exchange_api_endpoint"),
			ChainID:             "888",
		}
		cfg.DB = DBConfig{
			Host:     viper.GetString("testnet.database.host"),
			Port:     viper.GetString("testnet.database.port"),
			User:     viper.GetString("testnet.database.user"),
			Password: viper.GetString("testnet.database.password"),
			Table:    viper.GetString("testnet.database.table"),
		}

	default:
		log.Fatal("active can be only chain-id 888 (testnet).")
	}

	return cfg
}
