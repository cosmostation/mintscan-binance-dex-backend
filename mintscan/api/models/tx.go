package models

import (
	"encoding/json"
	"time"
)

// Txs is transaction data for result block
type Txs struct {
	Height    int64     `json:"height"`
	Result    bool      `json:"result"`
	TxHash    string    `json:"tx_hash"`
	Messages  []Message `json:"messages"`
	Memo      string    `json:"memo"`
	Code      uint32    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

// Message respresents tx meesage
type Message struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

// Message respresents tx signature
type Signature struct {
	Pubkey        string `json:"pubkey"`
	Address       string `json:"address"`
	Sequence      string `json:"sequence"`
	Signature     string `json:"signature"`
	AccountNumber string `json:"account_number"`
}

// TxRequestPayload respresents the data when
// receiving transaction request by its type and date
type TxRequestPayload struct {
	TxType    string `json:"tx_type"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// ValidatorMsgType verifies transaction message type
func ValidatorMsgType(msgType string) bool {
	msgTypes := []struct {
		msg, name string
	}{
		{"CreateOrderMsg", "dex/NewOrder"},
		{"CancelOrderMsg", "dex/CancelOrder"},
		{"TokenIssueMsg", "tokens/IssueMsg"},
		{"TokenBurnMsg", "tokens/BurnMsg"},
		{"TimeLockMsg", "tokens/TimeLockMsg"},
		{"TokenFreezeMsg", "tokens/FreezeMsg"},
		{"TokenUnfreezeMsg", "tokens/UnfreezeMsg"},
		{"TimeUnlockMsg", "tokens/TimeUnlockMsg"},
		{"TimeRelockMsg", "tokens/TimeRelockMsg"},
		{"HTLTMsg", "tokens/HTLTMsg"},
		{"DepositHTLTMsg", "tokens/DepositHTLTMsg"},
		{"ClaimHTLTMsg  ", "tokens/ClaimHTLTMsg"},
		{"RefundHTLTMsg ", "tokens/RefundHTLTMsg"},
		{"DexListMsg", "dex/ListMsg"},
		{"MintMsg   ", "tokens/MintMsg"},
		{"SendMsg", "cosmos-sdk/Send"},
		{"SubmitProposalMsg", "cosmos-sdk/MsgSubmitProposal"},
		{"DepositMsg       ", "cosmos-sdk/MsgDeposit"},
		{"VoteMsg          ", "cosmos-sdk/MsgVote"},
		{"SetAccountFlagsMsg", "scripts/SetAccountFlagsMsg"},
		{"MsgCreateValidator        ", "cosmos-sdk/MsgCreateValidator"},
		{"MsgRemoveValidator        ", "cosmos-sdk/MsgRemoveValidator"},
		{"MsgCreateValidatorProposal", "cosmos-sdk/MsgCreateValidatorProposal"},
	}

	for _, mt := range msgTypes {
		if msgType == mt.name {
			return true
		}
	}

	return false
}
