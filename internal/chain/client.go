package chain

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client interface {
	GetBlockByNumber(number uint64) (blk *Block, err error)
	GetLatestPeriod() (uint64, error)
	TraceBlockTransactions(number uint64) (traces []TransactionTrace, err error)
	GetTransactionByHash(hash string) (trx Transaction, err error)
	GetPeriodTransactions(period uint64) (trxs []Transaction, err error)
	GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags PbftBlockWithDags, err error)
	GetDagBlockByHash(hash string) (dag DagBlock, err error)
	GetPeriodDagBlocks(period uint64) (dags []DagBlock, err error)
	GetPreviousBlockCertVotes(period uint64) (vr VotesResponse, err error)
	GetValidatorsAtBlock(block_num uint64) (validators []Validator, err error)
	GetTotalAmountDelegated(block_num uint64) (totalAmountDelegated *big.Int, err error)
	GetTotalSupply(block_num uint64) (totalAmountDelegated *big.Int, err error)
	GetVersion() (version string, err error)
	GetGenesis() (genesis GenesisObject, err error)
	GetChainStats() (ns storage.FinalizationData, err error)
	SubscribeNewHeads() (chan Block, *rpc.ClientSubscription, error)
	GetBalanceAtBlock(address string, blockNumber uint64) (balance string, err error)
	GetLogs(fromBlock, toBlock uint64, addresses []string, topics [][]string) (logs []EventLog, err error)
	// Close disconnects from the node
	Close()
}
