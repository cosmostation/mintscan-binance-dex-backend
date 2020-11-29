module github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan

go 1.15

require (
	github.com/InjectiveLabs/sdk-go v1.11.6
	github.com/go-pg/pg v8.0.6+incompatible
	github.com/go-resty/resty/v2 v2.2.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc6
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	github.com/xlab/suplog v1.0.0
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/cosmos/cosmos-sdk => github.com/InjectiveLabs/cosmos-sdk v0.40.0-fix8

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
