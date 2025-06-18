package indexer

import (
	"encoding/json"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func TestLogsParsing(t *testing.T) {
	transaction_json := `{
		"blockHash": "0x055050d114dfa2befd5eef47ea415aa7baeb916bdd25c7794df1016bbbe71499",
    	"blockNumber": "0x1c12b3",
    	"contractAddress": null,
    	"cumulativeGasUsed": "0x146e0",
    	"from": "0xfc43217e71ec0a1cc480f3d210cd07cbde7374ec",
    	"gasUsed": "0xf4d8",
    	"logs": [
    	  	{
    	    	"address": "0x00000000000000000000000000000000000000fe",
    	    	"blockHash": "0x055050d114dfa2befd5eef47ea415aa7baeb916bdd25c7794df1016bbbe71499",
    	    	"blockNumber": "0x1c12b3",
    	    	"data": "0x0000000000000000000000000000000000000000000000000cc505042c728f8f",
    	    	"logIndex": "0x0",
    	    	"removed": false,
    	    	"topics": [
    	    		"0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7",
    	    		"0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec",
    	    		"0x000000000000000000000000e50b5452b2e8435404dbe06e6a05410c47b7583d"
    	    	],
    	    	"transactionHash": "0x689811a0705b89add2cd02d8a713bbd43c31c5afc123aeaca264494b375d6968",
    	    	"transactionIndex": "0x1"
    	  	},
    	  	{
    	   		"address": "0x00000000000000000000000000000000000000fe",
    	   		"blockHash": "0x055050d114dfa2befd5eef47ea415aa7baeb916bdd25c7794df1016bbbe71499",
    	   		"blockNumber": "0x1c12b3",
    	   		"data": "0x000000000000000000000000000000000000000000000088fd54c8913c769460",
    	   		"logIndex": "0x1",
    	   		"removed": false,
    	   		"topics": [
    	   			"0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b",
    	   			"0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec",
    	   			"0x000000000000000000000000e50b5452b2e8435404dbe06e6a05410c47b7583d"
    	   		],
    	   		"transactionHash": "0x689811a0705b89add2cd02d8a713bbd43c31c5afc123aeaca264494b375d6968",
    	   		"transactionIndex": "0x1"
    	  	}
    	],
    	"logsBloom": "0x8000000000000000080000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000010000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000008000000000000000040000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000800000000080000000000000000000000000000000000000020000000000000000000000400000000000000000000000000",
    	"status": "0x1",
    	"to": "0x00000000000000000000000000000000000000fe",
    	"transactionHash": "0x689811a0705b89add2cd02d8a713bbd43c31c5afc123aeaca264494b375d6968",
    	"transactionIndex": "0x1"
	}`

	logs := []models.EventLog{{

		Address:  "0x00000000000000000000000000000000000000fe",
		Data:     "0x0000000000000000000000000000000000000000000000000cc505042c728f8f",
		LogIndex: 0,
		Removed:  false,
		Topics: []string{
			"0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7",
			"0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec",
			"0x000000000000000000000000e50b5452b2e8435404dbe06e6a05410c47b7583d",
		},
		TransactionHash:  "0x689811a0705b89add2cd02d8a713bbd43c31c5afc123aeaca264494b375d6968",
		TransactionIndex: 1,
	},
		{
			Address:  "0x00000000000000000000000000000000000000fe",
			Data:     "0x000000000000000000000000000000000000000000000000a7a44a964be1f30a",
			LogIndex: 1,
			Removed:  false,
			Topics: []string{
				"0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b",
				"0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec",
				"0x000000000000000000000000e50b5452b2e8435404dbe06e6a05410c47b7583d",
			},
			TransactionHash:  "0x689811a0705b89add2cd02d8a713bbd43c31c5afc123aeaca264494b375d6968",
			TransactionIndex: 105,
		},
	}

	var trx common.Transaction
	_ = json.Unmarshal([]byte(transaction_json), &trx)

	parsedLogs := trx.ExtractLogs()

	for i, log := range parsedLogs {
		assert.Equal(t, log.Address, parsedLogs[i].Address)
		assert.Equal(t, log.Data, parsedLogs[i].Data)
		assert.Equal(t, logs[i].LogIndex, parsedLogs[i].LogIndex)
		assert.Equal(t, log.Removed, parsedLogs[i].Removed)
		assert.Equal(t, len(log.Topics), len(parsedLogs[i].Topics))
		assert.Equal(t, log.TransactionHash, parsedLogs[i].TransactionHash)
		assert.Equal(t, log.TransactionIndex, parsedLogs[i].TransactionIndex)
	}
}
