package models

type PbftResponse struct {
	Hash             Hash      `json:"hash"`
	Number           Counter   `json:"number"`
	Timestamp        Timestamp `json:"timestamp"`
	TransactionCount Counter   `json:"transactionCount"`
}
