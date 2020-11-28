package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config wraps all config.
type Config struct {
	Node   NodeConfig   `mapstructure:"node"`
	DB     DBConfig     `mapstructure:"database"`
	Web    WebConfig    `mapstructure:"web"`
	Market MarketConfig `mapstructure:"market"`
}

// NodeConfig wraps all node endpoints that are used in this project.
type NodeConfig struct {
	RPCNode             string `mapstructure:"rpc_node"`
	ExchangeAPIEndpoint string `mapstructure:"exchange_api_endpoint"`
	ChainID             string `mapstructure:"chain_id"`
}

// DBConfig wraps all required parameters for database connection.
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"dbname"`
}

// WebConfig wraps all required paramaters for boostraping web server.
type WebConfig struct {
	Port string `mapstructure:"port"`
}

// MarketConfig wraps all required params for market endpoints.
type MarketConfig struct {
	CoinGeckoEndpoint string `mapstructure:"coingecko_endpoint"`
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
