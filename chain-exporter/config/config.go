package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config defines all necessary parameters.
type Config struct {
	Node       NodeConfig       `mapstructure:"node"`
	DB         DBConfig         `mapstructure:"database"`
	Processing ProcessingConfig `mapstructure:"processing"`
}

type ProcessingConfig struct {
	GenesisHeight int  `mapstructure:"genesis_height"`
	IgnoreLogs    bool `mapstructure:"ignore_logs"`
}

// NodeConfig wraps all node endpoints that are used in this project.
type NodeConfig struct {
	RPCNode  string `mapstructure:"rpc_node"`
	GRPCNode string `mapstructure:"grpc_node"`
	ChainID  string `mapstructure:"chain_id"`
}

// DBConfig wraps all required parameters for database connection.
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"dbname"`
}

// ParseConfig attempts to read and parse config.yaml from the given path
// An error reading or parsing the config results in a panic.
func ParseConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../") // for test cases

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	chainID := viper.GetString("chain_id")
	if chainID == "" {
		log.Fatal("define active chain_id param in your config file.")
	}

	var config Config
	sub := viper.Sub(chainID)
	sub.Unmarshal(&config)

	if chainID == "888" {
		config.Node.ChainID = chainID
	} else {
		panic(chainID + " - chain not suppported yet")
	}

	return &config
}
