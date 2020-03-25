package codec

import (
	amino "github.com/tendermint/go-amino"

	"github.com/binance-chain/go-sdk/types"
)

// Codec is amino codec to serialize Binance Chain interfaces and data
var Codec *amino.Codec

// initializes upon package loading
func init() {
	Codec = types.NewCodec()
}
