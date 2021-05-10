module github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter

go 1.15

require (
	github.com/InjectiveLabs/sdk-go v1.20.3
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/ethereum/go-ethereum v1.9.25 // indirect
	github.com/go-pg/pg v8.0.6+incompatible
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	github.com/xlab/suplog v1.3.0
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
