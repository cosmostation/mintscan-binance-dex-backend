package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	Node NodeConfig     `toml:"node"`
	DB   DatabaseConfig `toml:"database"`
}

// NodeConfig defines endpoints for both RPC node and LCD REST API server
type NodeConfig struct {
	RPCEndpoint string `toml:"rpc_endpoint"`
	LCDEndpoint string `toml:"lcd_endpoint"`
}

// DatabaseConfig defines all database connection configuration parameters.
type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     uint64 `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	SSLMode  string `toml:"ssl_mode"`
}

// ParseConfig attempts to read and parse chain-exporter config from the given configPath.
// An error reading or parsing the config results in a panic.
func ParseConfig(configPath string) Config {
	if configPath == "" {
		log.Fatal("invalid configuration file")
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config"))
	}

	var cfg Config
	if _, err := toml.Decode(string(configData), &cfg); err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode config"))
	}

	return cfg
}
