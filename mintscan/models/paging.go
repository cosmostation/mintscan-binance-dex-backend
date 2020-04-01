package models

// Paging defines the structure for required params for handling pagination
type Paging struct {
	Total  int32 `json:"total"`  // total number of txs saved in database
	Before int32 `json:"before"` // can be either block height or index num
	After  int32 `json:"after"`  // can be either block height or index num
}
