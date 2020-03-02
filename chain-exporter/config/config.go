package config

import (
	"log"

	"github.com/pkg/errors"

	"github.com/spf13/viper"

	cmtypes "github.com/binance-chain/go-sdk/common/types"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	Node   NodeConfig   `yaml:"node"`
	DB     DBConfig     `yaml:"database"`
	Market MarketConfig `yaml:"market"`
}

// NodeConfig defines endpoints for both RPC node and LCD REST API server
type NodeConfig struct {
	RPCNode           string               `yaml:"rpc_node"`
	AcceleratedNode   string               `yaml:"accelerated_node"`
	APIServerEndpoint string               `yaml:"api_server_endpoint"`
	NetworkType       cmtypes.ChainNetwork `yaml:"network_type"`
}

// DBConfig defines all database connection configuration parameters.
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Table    string `yaml:"table"`
}

// MarketConfig defines endpoints where you parse market data from
type MarketConfig struct {
	CoinGeckoEndpoint string `yaml:"coingecko_endpoint"`
}

// ParseConfig attempts to read and parse chain-exporter config from the given configPath.
// An error reading or parsing the config results in a panic.
func ParseConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../") // for test cases
	viper.AddConfigPath("/home/ubuntu/mintscan-binance-dex-backend/chain-exporter/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config"))
	}

	cfg := Config{}

	if viper.GetString("active") == "" {
		log.Fatal("define active param in your config file.")
	}

	switch viper.GetString("active") {
	case "mainnet":
		cfg.Node = NodeConfig{
			RPCNode:           viper.GetString("mainnet.node.rpc_node"),
			AcceleratedNode:   viper.GetString("mainnet.node.accelerated_node"),
			APIServerEndpoint: viper.GetString("mainnet.node.api_server_endpoint"),
			NetworkType:       cmtypes.ProdNetwork,
		}
		cfg.DB = DBConfig{
			Host:     viper.GetString("mainnet.database.host"),
			Port:     viper.GetString("mainnet.database.port"),
			User:     viper.GetString("mainnet.database.user"),
			Password: viper.GetString("mainnet.database.password"),
			Table:    viper.GetString("mainnet.database.table"),
		}
		cfg.Market = MarketConfig{
			viper.GetString("mainnet.market.coingecko_endpoint"),
		}

	case "testnet":
		cfg.Node = NodeConfig{
			RPCNode:           viper.GetString("testnet.node.rpc_node"),
			AcceleratedNode:   viper.GetString("testnet.node.accelerated_node"),
			APIServerEndpoint: viper.GetString("testnet.node.api_server_endpoint"),
			NetworkType:       cmtypes.TestNetwork,
		}
		cfg.DB = DBConfig{
			Host:     viper.GetString("testnet.database.host"),
			Port:     viper.GetString("testnet.database.port"),
			User:     viper.GetString("testnet.database.user"),
			Password: viper.GetString("testnet.database.password"),
			Table:    viper.GetString("testnet.database.table"),
		}
		cfg.Market = MarketConfig{
			viper.GetString("testnet.market.coingecko_endpoint"),
		}

	default:
		log.Fatal("active can be either mainnet or testnet.")
	}

	return cfg
}
