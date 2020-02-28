module github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter

go 1.13

require (
	github.com/binance-chain/go-sdk v1.2.1
	github.com/dgraph-io/badger v1.6.0 // indirect
	github.com/fletaio/common v0.0.0-20190611002515-b9de21aaa2ad // indirect
	github.com/fletaio/core v0.0.0-20190618060323-4701ab309562
	github.com/fletaio/framework v0.0.0-20190611001721-27cef4d23774 // indirect
	github.com/fletaio/network v0.0.0-20190401100130-91d2fc4f149e // indirect
	github.com/go-pg/pg v8.0.6+incompatible
	github.com/go-resty/resty/v2 v2.2.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mr-tron/base58 v1.1.3 // indirect
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.8
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1
