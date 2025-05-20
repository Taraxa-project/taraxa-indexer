package common

import (
	"encoding/json"
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

const emptyInput = "0x"
const emptyReceiver = ""

func GetInternalTransactionTarget(trace TraceEntry) string {
	if trace.Action.To != "" {
		return trace.Action.To
	}
	return trace.Result.Address
}

func GetTransactionType(to, input, txType string, internal bool) models.TransactionType {
	if internal {
		if txType == "create" {
			return models.InternalContractCreation
		} else if txType == "call" && input != emptyInput {
			return models.InternalContractCall
		}
		return models.InternalTransfer
	} else {
		if to == emptyReceiver {
			return models.ContractCreation
		} else if input != emptyInput {
			return models.ContractCall
		}
		return models.Transfer
	}
}

type EventLog struct {
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockNumber      string   `json:"blockNumber"`
}

type Transaction struct {
	models.Transaction
	Logs             []EventLog     `json:"logs"`
	Nonce            models.Counter `json:"nonce"`
	GasPrice         models.Counter `json:"gasPrice"`
	GasUsed          models.Counter `json:"gasUsed"`
	TransactionIndex models.Counter `json:"transactionIndex"`
	ContractAddress  string         `json:"contractAddress"`
}

func (t *Transaction) SetTimestamp(timestamp uint64) {
	t.Transaction.Timestamp = timestamp
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	var rawStruct struct {
		BlockNumber string `json:"blockNumber"`
		From        string `json:"from"`
		Hash        string `json:"hash"`
		Input       string `json:"input"`
		Status      string `json:"status"`
		Timestamp   string `json:"timestamp"`
		To          string `json:"to"`
		Type        string `json:"type"`
		Value       string `json:"value"`

		Logs             []EventLog `json:"logs"`
		Nonce            string     `json:"nonce"`
		GasPrice         string     `json:"gasPrice"`
		GasUsed          string     `json:"gasUsed"`
		TransactionIndex string     `json:"transactionIndex"`

		ContractAddress string `json:"contractAddress"`
	}
	if err := json.Unmarshal(data, &rawStruct); err != nil {
		return err
	}

	t.Logs = rawStruct.Logs
	t.Nonce = ParseUInt(rawStruct.Nonce)
	t.GasPrice = ParseUInt(rawStruct.GasPrice)
	t.GasUsed = ParseUInt(rawStruct.GasUsed)
	t.TransactionIndex = ParseUInt(rawStruct.TransactionIndex)
	t.ContractAddress = rawStruct.ContractAddress

	t.Hash = rawStruct.Hash
	t.BlockNumber = ParseUInt(rawStruct.BlockNumber)
	t.From = rawStruct.From
	t.GasCost = t.GetFee().Uint64()
	t.Input = rawStruct.Input
	t.Status = ParseBool(rawStruct.Status)
	t.Timestamp = ParseUInt(rawStruct.Timestamp)
	t.To = rawStruct.To
	t.Value = rawStruct.Value

	t.Type = GetTransactionType(t.Transaction.To, t.Input, "", false)
	if t.Type == models.ContractCreation {
		t.To = t.ContractAddress
	}

	return nil
}

func (b *Transaction) GetModel() (trx *models.Transaction) {
	return &b.Transaction
}

func (t *Transaction) GetFee() *big.Int {
	return big.NewInt(0).Mul(big.NewInt(0).SetUint64(t.GasUsed), big.NewInt(0).SetUint64(t.GasPrice))
}

func (t *Transaction) ExtractLogs() (logs []models.EventLog) {
	for _, log := range t.Logs {
		eLog := models.EventLog{
			Address:          log.Address,
			Data:             log.Data,
			LogIndex:         ParseUInt(log.LogIndex),
			Name:             "",
			Params:           []string{},
			Removed:          log.Removed,
			Topics:           log.Topics,
			TransactionHash:  log.TransactionHash,
			TransactionIndex: ParseUInt(log.TransactionIndex),
		}
		logs = append(logs, eLog)
	}
	return logs
}
