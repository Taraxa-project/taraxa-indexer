package common

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
)

type CallData struct {
	Name   string `json:"name"`
	Params any    `json:"params"`
}

func splitFunctionIDFromData(data []byte) ([]byte, []byte, error) {
	if len(data) < 4 {
		return nil, nil, fmt.Errorf("transaction data is too short")
	}
	return data[:4], data[4:], nil
}

func DecodeTransaction(tx models.Transaction) (functionSig string, params any, err error) {
	if tx.Input == "" {
		return
	}
	relevantAbi := contracts.ContractABIs[strings.ToLower(tx.To)]
	if relevantAbi == "" {
		return
	}
	contractABI, err := abi.JSON(strings.NewReader(relevantAbi))
	if err != nil {
		return
	}

	trimmed := strings.TrimPrefix(tx.Input, "0x")
	bytes, err := hex.DecodeString(trimmed)

	if err != nil {
		return
	}

	funcId, data, err := splitFunctionIDFromData(bytes)
	if err != nil {
		return
	}
	// Decode the transaction
	method, err := contractABI.MethodById(funcId)

	if method == nil {
		fmt.Println(method, err)
	}

	functionSig = method.Sig

	if err != nil || method == nil {
		return
	}

	unpacked, err := unpackParams(contractABI, method, data)

	if err != nil {
		return
	}

	params, err = ParseToString(unpacked)

	if err != nil {
		return
	}

	return
}

func unpackParams(contractABI abi.ABI, method *abi.Method, data []byte) ([]interface{}, error) {
	var args abi.Arguments
	if method, ok := contractABI.Methods[method.Name]; ok {
		if len(data)%32 != 0 {
			return nil, fmt.Errorf("abi: improperly formatted output: %x", data)
		}
		args = method.Inputs
	}
	unpacked, err := args.Unpack(data)
	return unpacked, err
}

func ProcessTransaction(trx *models.Transaction) (err error) {
	sig, params, err := DecodeTransaction(*trx)

	if sig == "" && params != nil {
		return
	}

	if err != nil {
		log.WithError(err).WithFields(log.Fields{"hash": trx.Hash}).Debug("DecodeTransaction error")
		return
	}

	trx.Calldata = &models.CallData{
		Name:   sig,
		Params: params,
	}

	return
}
