package indexer

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func splitFunctionIDFromData(data []byte) ([]byte, []byte, error) {
	if len(data) < 4 {
		return nil, nil, fmt.Errorf("transaction data is too short")
	}
	return data[:4], data[4:], nil
}

func DecodeTransaction(tx chain.Transaction) (functionSig string, params []string, err error) {
	if tx.Data == "" {
		return "", nil, nil
	}
	relevantAbi := contracts.ContractABIs[strings.ToLower(tx.To)]
	if relevantAbi == "" {
		return "", nil, nil
	}
	contractABI, error := abi.JSON(strings.NewReader(relevantAbi))
	if error != nil {
		return "", nil, error
	}

	trimmed := strings.TrimPrefix(tx.Data, "0x")
	bytes, err := hex.DecodeString(trimmed)

	if err != nil {
		return "", nil, err
	}

	funcId, data, err := splitFunctionIDFromData(bytes)
	if err != nil {
		return "", nil, err
	}
	// Decode the transaction
	method, err := contractABI.MethodById(funcId)

	functionSig = method.Sig

	if err != nil {
		return "", nil, err
	}

	if method == nil {
		return
	}

	// TODO: move to separate function
	var args abi.Arguments
	if method, ok := contractABI.Methods[method.Name]; ok {
		if len(data)%32 != 0 {
			log.Fatal("failed to decode transaction")
			// return nil, fmt.Errorf("abi: improperly formatted output: %q - Bytes: %+v", data, data)
		}
		args = method.Inputs
	}
	unpacked, err := args.Unpack(data)
	// END TODO

	if err != nil {
		return "", nil, err
	}

	params, err = common.ParseToStringSlice(unpacked)

	if err != nil {
		return "", nil, err
	}

	return functionSig, params, nil
}
