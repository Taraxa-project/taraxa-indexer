package indexer

import (
	"encoding/hex"
	"fmt"
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
	relevantAbi := contracts.ContractABIs[tx.To]
	if relevantAbi == "" {
		return "", nil, nil
	}
	contractABI, error := abi.JSON(strings.NewReader(relevantAbi))
	if error != nil {
		return "", nil, error
	}

	trimmed := strings.TrimPrefix(tx.Data, "0x")
	bytes, err := hex.DecodeString(trimmed)
	funcId, data, err := splitFunctionIDFromData(bytes)
	if err != nil {
		return "", nil, err
	}
	// Decode the transaction
	method, err := contractABI.MethodById(funcId)

	functionSig = method.Sig

	fmt.Println("DecodeTransaction Method: ", functionSig)

	if err != nil {
		return "", nil, err
	}

	if method == nil {
		return
	}

	fmt.Println("Len of data: ", len(data))

	fmt.Println(method.Inputs)

	unpacked, err := contractABI.Unpack(method.Name, data)

	fmt.Println("DecodeTransaction Unpacked: ", unpacked)
	if err != nil {
		return "", nil, err
	}

	params, err = common.ParseToStringSlice(unpacked)

	if err != nil {
		return "", nil, err
	}

	return functionSig, params, nil

}
