module github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter

go 1.13

require (
	github.com/binance-chain/go-sdk v1.2.1
	github.com/go-pg/pg v8.0.6+incompatible
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.0.0
	github.com/tendermint/tendermint v0.32.3
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1
