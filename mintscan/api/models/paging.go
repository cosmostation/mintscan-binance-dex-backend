package models

// Paging wraps required params for handling pagination
type Paging struct {
	Total  int   `json:"total"` // total number of txs saved in database
	Before int32 `json:"before"`
	After  int32 `json:"after"`
}
