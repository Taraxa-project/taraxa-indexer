package indexer

import (
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func TestTraceInternalCreationParsing(t *testing.T) {
	transaction_hash := "0x516a20b6ee0c19eb8188ff45176e1cdb6230aa740cbdce478f7856bed89aa2f0"
	transaction_json := `{
		"blockHash": "0x5d189eb4214bff0bb81b3564742cd040ed38fb3306e05923cbd1006c37a6cb53",
		"blockNumber": "0x5487c",
		"from": "0xa44ef3fee7598e86f2c6dfbf1c38e64f2d55f6e7",
		"gas": "0x5b8d80",
		"gasPrice": "0x0",
		"gasUsed": "0x5af868",
		"hash": "0x516a20b6ee0c19eb8188ff45176e1cdb6230aa740cbdce478f7856bed89aa2f0",
		"input": "0x7101f091000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000003b9aca000000000000000000000000000000000000000000000000000000000000000005575441524100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000055754415241000000000000000000000000000000000000000000000000000000",
		"nonce": "0x0",
		"r": "0xe41fe3baf6ef792fa3d33af556767c84528ee3b8387fcd94d33fb3f12bff2d03",
		"s": "0x4d4462271713d1b1808368952cf44bef70fdd8304246cf5ee0ffa38d771b55d9",
		"to": "0x26d80dd41f67a8c97d5e0a233e8f7975cfaa23ed",
		"transactionIndex": "0x0",
		"v": "0x0",
		"value": "0x0"
	}`
	traces_json := `[
		{
		  "output": "0x0000000000000000000000003a68620feeca3712d0b420a6ef41e4fa0683b18b",
		  "stateDiff": null,
		  "trace": [
			{
			  "action": {
				"callType": "call",
				"from": "0xa44ef3fee7598e86f2c6dfbf1c38e64f2d55f6e7",
				"gas": "0x7d845",
				"input": "0x7101f091000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000003b9aca000000000000000000000000000000000000000000000000000000000000000005575441524100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000055754415241000000000000000000000000000000000000000000000000000000",
				"to": "0x26d80dd41f67a8c97d5e0a233e8f7975cfaa23ed",
				"value": "0x0"
			  },
			  "result": {
				"gasUsed": "0x79d30",
				"output": "0x0000000000000000000000003a68620feeca3712d0b420a6ef41e4fa0683b18b"
			  },
			  "subtraces": 1,
			  "traceAddress": [],
			  "type": "call"
			},
			{
			  "action": {
				"from": "0x26d80dd41f67a8c97d5e0a233e8f7975cfaa23ed",
				"gas": "0x73303",
				"init": "0x608060405234801562000010575f80fd5b5060405162000df738038062000df7833981810160405281019062000036919062000320565b845f908162000046919062000610565b50836001908162000058919062000610565b508260025f6101000a81548160ff021916908360ff160217905550816003819055508160045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505050505050620006f4565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200012882620000e0565b810181811067ffffffffffffffff821117156200014a5762000149620000f0565b5b80604052505050565b5f6200015e620000c7565b90506200016c82826200011d565b919050565b5f67ffffffffffffffff8211156200018e576200018d620000f0565b5b6200019982620000e0565b9050602081019050919050565b5f5b83811015620001c5578082015181840152602081019050620001a8565b5f8484015250505050565b5f620001e6620001e08462000171565b62000153565b905082815260208101848484011115620002055762000204620000dc565b5b62000212848285620001a6565b509392505050565b5f82601f830112620002315762000230620000d8565b5b815162000243848260208601620001d0565b91505092915050565b5f60ff82169050919050565b62000263816200024c565b81146200026e575f80fd5b50565b5f81519050620002818162000258565b92915050565b5f819050919050565b6200029b8162000287565b8114620002a6575f80fd5b50565b5f81519050620002b98162000290565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620002ea82620002bf565b9050919050565b620002fc81620002de565b811462000307575f80fd5b50565b5f815190506200031a81620002f1565b92915050565b5f805f805f60a086880312156200033c576200033b620000d0565b5b5f86015167ffffffffffffffff8111156200035c576200035b620000d4565b5b6200036a888289016200021a565b955050602086015167ffffffffffffffff8111156200038e576200038d620000d4565b5b6200039c888289016200021a565b9450506040620003af8882890162000271565b9350506060620003c288828901620002a9565b9250506080620003d5888289016200030a565b9150509295509295909350565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200043157607f821691505b602082108103620004475762000446620003ec565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004ab7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200046e565b620004b786836200046e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f620004f8620004f2620004ec8462000287565b620004cf565b62000287565b9050919050565b5f819050919050565b6200051383620004d8565b6200052b6200052282620004ff565b8484546200047a565b825550505050565b5f90565b6200054162000533565b6200054e81848462000508565b505050565b5b818110156200057557620005695f8262000537565b60018101905062000554565b5050565b601f821115620005c4576200058e816200044d565b62000599846200045f565b81016020851015620005a9578190505b620005c1620005b8856200045f565b83018262000553565b50505b505050565b5f82821c905092915050565b5f620005e65f1984600802620005c9565b1980831691505092915050565b5f620006008383620005d5565b9150826002028217905092915050565b6200061b82620003e2565b67ffffffffffffffff811115620006375762000636620000f0565b5b62000643825462000419565b6200065082828562000579565b5f60209050601f83116001811462000686575f841562000671578287015190505b6200067d8582620005f3565b865550620006ec565b601f19841662000696866200044d565b5f5b82811015620006bf5784890151825560018201915060208501945060208101905062000698565b86831015620006df5784890151620006db601f891682620005d5565b8355505b6001600288020188555050505b505050505050565b6106f580620007025f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c806306fdde031461006457806318160ddd14610082578063313ce567146100a057806370a08231146100be57806395d89b41146100ee578063a9059cbb1461010c575b5f80fd5b61006c610128565b6040516100799190610459565b60405180910390f35b61008a6101b3565b6040516100979190610491565b60405180910390f35b6100a86101b9565b6040516100b591906104c5565b60405180910390f35b6100d860048036038101906100d3919061053c565b6101cb565b6040516100e59190610491565b60405180910390f35b6100f66101e0565b6040516101039190610459565b60405180910390f35b61012660048036038101906101219190610591565b61026c565b005b5f8054610134906105fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610160906105fc565b80156101ab5780601f10610182576101008083540402835291602001916101ab565b820191905f5260205f20905b81548152906001019060200180831161018e57829003601f168201915b505050505081565b60035481565b60025f9054906101000a900460ff1681565b6004602052805f5260405f205f915090505481565b600180546101ed906105fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610219906105fc565b80156102645780601f1061023b57610100808354040283529160200191610264565b820191905f5260205f20905b81548152906001019060200180831161024757829003601f168201915b505050505081565b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101580156102b857505f81115b6102c0575f80fd5b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461030c9190610659565b925050819055508060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461035f919061068c565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516103c39190610491565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156104065780820151818401526020810190506103eb565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61042b826103cf565b61043581856103d9565b93506104458185602086016103e9565b61044e81610411565b840191505092915050565b5f6020820190508181035f8301526104718184610421565b905092915050565b5f819050919050565b61048b81610479565b82525050565b5f6020820190506104a45f830184610482565b92915050565b5f60ff82169050919050565b6104bf816104aa565b82525050565b5f6020820190506104d85f8301846104b6565b92915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61050b826104e2565b9050919050565b61051b81610501565b8114610525575f80fd5b50565b5f8135905061053681610512565b92915050565b5f60208284031215610551576105506104de565b5b5f61055e84828501610528565b91505092915050565b61057081610479565b811461057a575f80fd5b50565b5f8135905061058b81610567565b92915050565b5f80604083850312156105a7576105a66104de565b5b5f6105b485828601610528565b92505060206105c58582860161057d565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061061357607f821691505b602082108103610626576106256105cf565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61066382610479565b915061066e83610479565b92508282039050818111156106865761068561062c565b5b92915050565b5f61069682610479565b91506106a183610479565b92508282019050808211156106b9576106b861062c565b5b9291505056fea2646970667358221220d589e68cb1f71824702aa7aadd3d7be708b7e646850136685ed32a80f2bbf12f64736f6c6343000815003300000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000003b9aca00000000000000000000000000a44ef3fee7598e86f2c6dfbf1c38e64f2d55f6e70000000000000000000000000000000000000000000000000000000000000005575441524100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000055754415241000000000000000000000000000000000000000000000000000000",
				"value": "0x0"
			  },
			  "result": {
				"address": "0x3a68620feeca3712d0b420a6ef41e4fa0683b18b",
				"code": "0x608060405234801561000f575f80fd5b5060043610610060575f3560e01c806306fdde031461006457806318160ddd14610082578063313ce567146100a057806370a08231146100be57806395d89b41146100ee578063a9059cbb1461010c575b5f80fd5b61006c610128565b6040516100799190610459565b60405180910390f35b61008a6101b3565b6040516100979190610491565b60405180910390f35b6100a86101b9565b6040516100b591906104c5565b60405180910390f35b6100d860048036038101906100d3919061053c565b6101cb565b6040516100e59190610491565b60405180910390f35b6100f66101e0565b6040516101039190610459565b60405180910390f35b61012660048036038101906101219190610591565b61026c565b005b5f8054610134906105fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610160906105fc565b80156101ab5780601f10610182576101008083540402835291602001916101ab565b820191905f5260205f20905b81548152906001019060200180831161018e57829003601f168201915b505050505081565b60035481565b60025f9054906101000a900460ff1681565b6004602052805f5260405f205f915090505481565b600180546101ed906105fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610219906105fc565b80156102645780601f1061023b57610100808354040283529160200191610264565b820191905f5260205f20905b81548152906001019060200180831161024757829003601f168201915b505050505081565b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101580156102b857505f81115b6102c0575f80fd5b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461030c9190610659565b925050819055508060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461035f919061068c565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516103c39190610491565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156104065780820151818401526020810190506103eb565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61042b826103cf565b61043581856103d9565b93506104458185602086016103e9565b61044e81610411565b840191505092915050565b5f6020820190508181035f8301526104718184610421565b905092915050565b5f819050919050565b61048b81610479565b82525050565b5f6020820190506104a45f830184610482565b92915050565b5f60ff82169050919050565b6104bf816104aa565b82525050565b5f6020820190506104d85f8301846104b6565b92915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61050b826104e2565b9050919050565b61051b81610501565b8114610525575f80fd5b50565b5f8135905061053681610512565b92915050565b5f60208284031215610551576105506104de565b5b5f61055e84828501610528565b91505092915050565b61057081610479565b811461057a575f80fd5b50565b5f8135905061058b81610567565b92915050565b5f80604083850312156105a7576105a66104de565b5b5f6105b485828601610528565b92505060206105c58582860161057d565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061061357607f821691505b602082108103610626576106256105cf565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61066382610479565b915061066e83610479565b92508282039050818111156106865761068561062c565b5b92915050565b5f61069682610479565b91506106a183610479565b92508282019050808211156106b9576106b861062c565b5b9291505056fea2646970667358221220d589e68cb1f71824702aa7aadd3d7be708b7e646850136685ed32a80f2bbf12f64736f6c63430008150033",
				"gasUsed": "0x70f41"
			  },
			  "subtraces": 0,
			  "traceAddress": [
				0
			  ],
			  "type": "create"
			}
		  ],
		  "vmTrace": null
		}
	  ]`

	mc := chain.MakeMockClient()
	mc.AddTransactionFromJson(transaction_json)
	tt, _ := mc.GetTransactionByHash(transaction_hash)
	trx := tt.ToModelWithTimestamp(1)
	assert.Equal(t, transaction_hash, trx.Hash)
	assert.Equal(t, uint64(0x5487c), trx.BlockNumber)
	assert.Equal(t, models.ContractCall, trx.Type)

	bc := MakeTestBlockContext(mc, trx.BlockNumber)

	mc.AddTracesFromJson(transaction_hash, traces_json)

	transactions_trace, _ := bc.Client.TraceBlockTransactions(trx.BlockNumber)
	// Have one transaction with 2 internal transactions
	trx_count := 1
	internal_count := 1
	assert.Equal(t, trx_count, len(transactions_trace))
	assert.Equal(t, trx_count+internal_count, len(transactions_trace[0].Trace))

	err := bc.processTransactions([]string{trx.Hash})

	assert.Equal(t, err, nil)
	bc.commit()

	int_trx := bc.Storage.GetInternalTransactions(trx.Hash)
	assert.Equal(t, internal_count, len(int_trx.Data))

	// count transactions to compare it with db data
	addr_trx_count := make(map[string]int)
	for i, trace := range transactions_trace[0].Trace {
		if i == 0 {
			// skip first element as it is original transaction data
			continue
		}
		internal := int_trx.Data[i-1]
		assert.Equal(t, models.InternalContractCreation, internal.Type)
		assert.Equal(t, internal.Hash, trx.Hash)
		assert.Equal(t, internal.From, trace.Action.From)
		assert.Equal(t, internal.To, trace.Result.Address)
		assert.Equal(t, internal.Value, trace.Action.Value)
	}

	for addr, count := range addr_trx_count {
		res, _ := storage.GetObjectsPage[models.Transaction](bc.Storage, addr, 0, 20)
		assert.Equal(t, len(res), count, addr)
	}
}