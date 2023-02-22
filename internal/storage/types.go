package storage

// Address defines the model for an address aggregate.
type Address struct {
	Address   string `json:"address"`
	TxTotal   uint64 `json:"txTotal"`
	DagTotal  uint64 `json:"dagTotal"`
	PbftTotal uint64 `json:"pbftTotal"`
}
