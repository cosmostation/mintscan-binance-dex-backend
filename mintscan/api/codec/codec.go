package codec

import (
	amino "github.com/tendermint/go-amino"

	"github.com/binance-chain/go-sdk/types"
)

// Codec is amino codec to marshal/unmarshal Binance Chain interfaces and etc.
var Codec *amino.Codec

// initializes upon package loading
func init() {
	Codec = types.NewCodec()
}
