package pebble

import "github.com/Taraxa-project/taraxa-indexer/models"

type OldDag struct {
	Hash             models.Hash      `json:"hash"`
	Sender           models.Address   `json:"sender"`
	Level            models.Counter   `json:"level"`
	Timestamp        models.Timestamp `json:"timestamp"`
	TransactionCount models.Counter   `json:"transactionCount"`
}
