package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	cmtypes "github.com/binance-chain/go-sdk/common/types"
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
	RPCNode                string               `mapstructure:"rpc_node"`
	AcceleratedNode        string               `mapstructure:"accelerated_node"`
	APIServerEndpoint      string               `mapstructure:"api_server_endpoint"`
	ExplorerServerEndpoint string               `mapstructure:"explorer_server_endpoint"`
	NetworkType            cmtypes.ChainNetwork `mapstructure:"network_type"`
}

// DBConfig wraps all required parameters for database connection.
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Table    string `mapstructure:"table"`
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
	viper.AddConfigPath("../")                                                 // for test cases
	viper.AddConfigPath("/home/ubuntu/mintscan-binance-dex-backend/mintscan/") // for production

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	if viper.GetString("network_type") == "" {
		log.Fatal("define active param in your config file.")
	}

	var config Config
	sub := viper.Sub(viper.GetString("network_type"))
	sub.Unmarshal(&config)

	if viper.GetString("network_type") == "mainnet" {
		config.Node.NetworkType = cmtypes.ProdNetwork
	} else {
		config.Node.NetworkType = cmtypes.TestNetwork
	}

	return &config
}
