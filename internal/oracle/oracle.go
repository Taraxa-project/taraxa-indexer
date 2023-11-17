package oracle

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/contracts"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type ValidatorData struct {
	account   string
	rank      uint32
	rating    uint32
	apy       uint32
	fromBlock uint64
	toBlock   uint64
}

func updateOracle(client chain.WsClient, methodName string, yield storage.ValidatorYield) (result string, err error) {
	oracleAddress := "0x6d68DC5F5B30aFd0A4dF915d703529aB2970D443"

	contractAbi := contracts.ApyOracle

	parsedABI, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return "", err
	}

	validatorData := ValidatorData{
		account:   yield.Account,
		

	// Pack the method call with arguments
	data, err := parsedABI.Pack(methodName, )
	if err != nil {
		return "", err
	}

	// Prepare the call
	callMsg := map[string]interface{}{
		"to":   oracleAddress,
		"data": fmt.Sprintf("0x%x", data),
	}

	// The block number parameter can also be "latest" or "pending"
	var blockNumber = "latest"

	// Perform the call
	err = client.rpc.CallContext(context.Background(), &result, "eth_call", callMsg, blockNumber)
	if err != nil {
		return "", err
	}

	// The result is the return value of the contract method, hex encoded
	return result, nil

}

func PushOnChain(yields []storage.ValidatorYield, client chain.Client) {
	for _, yield := range yields {
		result, err := updateOracle(wsClient, "yourContractMethodName", yield)
		if err != nil {
			fmt.Printf("Failed to call contract method: %v\n", err)
			return
		}

		fmt.Printf("Result from contract call: %s\n", result)

		// If the result is not a simple value and you know the expected types, you can unpack it
		// For example, if the expected type is *big.Int
		var intValue *big.Int
		err = abi.U256(big.NewInt(0).SetBytes(common.FromHex(result)), &intValue)
		if err != nil {
			fmt.Printf("Failed to decode result: %v\n", err)
			return
		}

		fmt.Printf("Decoded result: %v\n", intValue)
	}
}
