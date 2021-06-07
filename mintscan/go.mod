module github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan

go 1.16

require (
	github.com/InjectiveLabs/sdk-go v1.21.3
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-pg/pg v8.0.6+incompatible
	github.com/go-resty/resty/v2 v2.2.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	github.com/xlab/suplog v1.3.0
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
