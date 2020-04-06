package models

type (
	// TxMsgFee defines the structure for transaction message type fees API
	TxMsgFee struct {
		MsgType           string         `json:"msg_type,omitempty"`
		Fee               int            `json:"fee,omitempty"`
		FeeFor            int            `json:"fee_for,omitempty"`
		FixedFeeParams    *FixedFeeParam `json:"fixed_fee_params,omitempty"`
		MultiTransferFee  int            `json:"multi_transfer_fee,omitempty"`
		LowerLimitAsMulti int            `json:"lower_limit_as_multi,omitempty"`
		DexFeeFields      []struct {
			FeeName  string `json:"fee_name,omitempty"`
			FeeValue int    `json:"fee_value,omitempty"`
		} `json:"dex_fee_fields,omitempty"`
	}

	// FixedFeeParam wraps fixed fee param
	FixedFeeParam struct {
		MsgType string `json:"msg_type,omitempty"`
		Fee     int    `json:"fee,omitempty"`
		FeeFor  int    `json:"fee_for,omitempty"`
	}
)

//
