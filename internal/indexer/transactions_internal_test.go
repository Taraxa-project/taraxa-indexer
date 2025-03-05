package indexer

import (
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func MakeTestBlockContext(mc *chain.ClientMock, blockNumber uint64) *blockContext {
	st := pebble.NewStorage("")
	bd, err := chain.GetBlockData(mc, blockNumber)
	if err != nil {
		panic(err)
	}
	bc := MakeBlockContext(st, mc, new(common.Config), storage.MakeAccountsMap())
	bc.SetBlockData(bd)

	return bc
}

func TestTraceParsing(t *testing.T) {
	transaction_hash := "0x1312631675d90f8625f6661bb384708d7b020f1894e12d34616c6258e6edf372"
	transaction_json := `{
		"blockHash": "0x7af7c63eb1caefe41e33630606d245be9a172892a796fffaa40c77e50e414405",
		"blockNumber": "0x5487c",
		"from": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
		"gas": "0x5b8d80",
		"gasPrice": "0x0",
		"gasUsed": "0x5af868",
		"hash": "0x1312631675d90f8625f6661bb384708d7b020f1894e12d34616c6258e6edf372",
		"input": "0xd2d745b10000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000000090000000000000000000000000dc0d841f962759da25547c686fa440cf6c28c61000000000000000000000000f1c587a22fbf80af80446fa17e7322952f18456c000000000000000000000000cddb0d484ca1c625ffca0882396ef34ffff242e300000000000000000000000010ce6f9c7c22f82214c40755b3eea5f126a7148d000000000000000000000000d42eaa28c5eafee9a0040a7ac74dd3f4b57678bd000000000000000000000000ec15db470db85cc75b0e3fa5b6a0c607a5e8c64a000000000000000000000000a4195def477491ef7f00b8688c9b8032cd71bb2a0000000000000000000000008f1567bb4381f4ed53dbeb3c0dca5c4f189a111000000000000000000000000052124d5982576507dd4a18d6607225e64be168bb0000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000",
		"nonce": "0x0",
		"r": "0xe41fe3baf6ef792fa3d33af556767c84528ee3b8387fcd94d33fb3f12bff2d03",
		"s": "0x4d4462271713d1b1808368952cf44bef70fdd8304246cf5ee0ffa38d771b55d9",
		"to": "0x1578f035581f664efa85a6da822464bd9edd8851",
		"transactionIndex": "0x0",
		"v": "0x0",
		"value": "0x82f79cd9000",
		"status": "0x1"
	}`
	traces_json := `[
	{
		"output": "0x",
		"stateDiff": null,
		"trace": [
			{
				"action": {
					"callType": "call",
					"from": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"gas": "0x5af868",
					"input": "0xd2d745b10000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000000090000000000000000000000000dc0d841f962759da25547c686fa440cf6c28c61000000000000000000000000f1c587a22fbf80af80446fa17e7322952f18456c000000000000000000000000cddb0d484ca1c625ffca0882396ef34ffff242e300000000000000000000000010ce6f9c7c22f82214c40755b3eea5f126a7148d000000000000000000000000d42eaa28c5eafee9a0040a7ac74dd3f4b57678bd000000000000000000000000ec15db470db85cc75b0e3fa5b6a0c607a5e8c64a000000000000000000000000a4195def477491ef7f00b8688c9b8032cd71bb2a0000000000000000000000008f1567bb4381f4ed53dbeb3c0dca5c4f189a111000000000000000000000000052124d5982576507dd4a18d6607225e64be168bb0000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000000000000000000000000000000000000000000000000000000000e8d4a51000",
					"to": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"value": "0x82f79cd9000"
				},
				"result": {
					"gasUsed": "0x43b61",
					"output": "0x"
				},
				"subtraces": 9,
				"traceAddress": [],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					0
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					1
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					2
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					3
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					4
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					5
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					6
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					7
				],
				"type": "call"
			},
			{
				"action": {
					"callType": "call",
					"from": "0x1578f035581f664efa85a6da822464bd9edd8851",
					"gas": "0x8fc",
					"input": "0x",
					"to": "0x99a2d5feaecb1a729d4f9af4197cc03bb9a37bc3",
					"value": "0xe8d4a51000"
				},
				"result": {
					"gasUsed": "0x0",
					"output": "0x"
				},
				"subtraces": 0,
				"traceAddress": [
					8
				],
				"type": "call"
			}
		],
		"vmTrace": null
	}]`

	mc := chain.MakeMockClient()
	mc.AddTransactionFromJson(transaction_json)
	trx, _ := mc.GetTransactionByHash(transaction_hash)
	trx.SetTimestamp(1)

	mc.AddTracesFromJson(transaction_hash, traces_json)

	assert.Equal(t, transaction_hash, trx.Hash)
	assert.Equal(t, uint64(0x5487c), trx.BlockNumber)
	assert.Equal(t, models.ContractCall, trx.Type)

	pbft := &chain.Block{}
	pbft.Number = trx.BlockNumber
	pbft.Transactions = []string{trx.Hash}
	pbft.TransactionCount = 1
	mc.AddPbftBlock(trx.BlockNumber, pbft)

	bc := MakeTestBlockContext(mc, trx.BlockNumber)

	transactions_trace, _ := bc.Client.TraceBlockTransactions(trx.BlockNumber)
	// Have one transaction with 9 internal transactions
	trx_count := 1
	internal_count := 9
	assert.Equal(t, trx_count, len(transactions_trace))
	assert.Equal(t, trx_count+internal_count, len(transactions_trace[0].Trace))

	err := bc.processTransactions()

	assert.Equal(t, err, nil)
	bc.commit()

	int_trx := bc.Storage.GetInternalTransactions(trx.Hash)
	assert.Equal(t, internal_count, len(int_trx.Data))

	// count transactions to compare it with db data
	addr_trx_count := make(map[string]int)
	for i, trace := range transactions_trace[0].Trace {
		addr_trx_count[trace.Action.From] += 1
		addr_trx_count[trace.Action.To] += 1
		if i == 0 {
			// skip first element as it is original transaction data
			continue
		}
		internal := int_trx.Data[i-1]
		assert.Equal(t, models.InternalTransfer, internal.Type)
		assert.Equal(t, internal.Hash, trx.Hash)
		assert.Equal(t, internal.From, trace.Action.From)
		assert.Equal(t, internal.To, trace.Action.To)
		assert.Equal(t, internal.Value, common.ParseStringToBigInt(trace.Action.Value))
	}

	for addr, count := range addr_trx_count {
		res, _ := storage.GetObjectsPage[storage.Transaction](bc.Storage, addr, 0, 20)
		assert.Equal(t, len(res), count, addr)
	}
}
