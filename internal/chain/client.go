package chain

import (
	"math/big"

	"github.com/Taraxa-project/taraxa-go-client/taraxa_client/dpos_contract_client/dpos_interface"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client interface {
	GetBlockByNumber(number uint64) (blk *Block, err error)
	GetLatestPeriod() (uint64, error)
	TraceBlockTransactions(number uint64) (traces []TransactionTrace, err error)
	GetTransactionByHash(hash string) (trx Transaction, err error)
	GetPeriodTransactions(period uint64) (trxs []Transaction, err error)
	GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags *PbftBlockWithDags, err error)
	GetDagBlockByHash(hash string) (dag *DagBlock, err error)
	GetPeriodDagBlocks(period uint64) (dags []DagBlock, err error)
	GetPreviousBlockCertVotes(period uint64) (vr *VotesResponse, err error)
	GetValidatorsAtBlock(block_num *big.Int) (validators []dpos_interface.DposInterfaceValidatorData, err error)
	GetGenesis() (genesis *GenesisObject, err error)
	GetChainStats() (ns *storage.FinalizationData, err error)
	SubscribeNewHeads() (chan *Block, *rpc.ClientSubscription, error)
	// Close disconnects from the node
	Close()
}
