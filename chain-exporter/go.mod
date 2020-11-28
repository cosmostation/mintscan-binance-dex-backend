module github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter

go 1.15

require (
	github.com/InjectiveLabs/sdk-go v1.11.0
	github.com/bugsnag/panicwrap v1.2.0 // indirect
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/go-pg/pg v8.0.6+incompatible
	github.com/go-resty/resty/v2 v2.2.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.0-rc6
	github.com/xlab/suplog v1.0.0
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/cosmos/cosmos-sdk => github.com/InjectiveLabs/cosmos-sdk v0.40.0-fix8

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
