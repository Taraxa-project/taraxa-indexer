package storage

import (
	"math/big"
	"strings"
	"sync"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/events"
	"github.com/Taraxa-project/taraxa-indexer/models"
	log "github.com/sirupsen/logrus"
)

// AddressStats defines the model for an address aggregate.
type AddressStats struct {
	models.StatsResponse
	Address string       `json:"address"`
	Balance *big.Int     `json:"balance"`
	mutex   sync.RWMutex `rlp:"-"`
}

func MakeEmptyAddressStats(addr string) *AddressStats {
	data := new(AddressStats)
	data.Address = addr
	data.Balance = big.NewInt(0)
	return data
}

func (a *AddressStats) RegisterValidatorBlock(blockHeight uint64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.ValidatorRegisteredBlock = &blockHeight
}

func (a *AddressStats) AddTransaction(trx *models.Transaction) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.TransactionsCount++
	if a.Address == trx.From {
		a.LastTransactionTimestamp = &trx.Timestamp
	}
	if a.Address == trx.To && (trx.Type == models.ContractCreation || trx.Type == models.InternalContractCreation) {
		a.ContractRegisteredTimestamp = &trx.Timestamp
	}
	return a.TransactionsCount
}

func (a *AddressStats) AddPbft(timestamp models.Timestamp) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.PbftCount++
	a.LastPbftTimestamp = &timestamp
	return a.PbftCount
}

func (a *AddressStats) AddDag(timestamp models.Timestamp) uint64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.DagsCount++
	a.LastDagTimestamp = &timestamp
	return a.DagsCount
}

func (a *AddressStats) IsEqual(b *AddressStats) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if a.Address == b.Address && a.TransactionsCount == b.TransactionsCount && a.DagsCount == b.DagsCount && a.PbftCount == b.PbftCount {
		return true
	}
	return false
}

func (a *AddressStats) IsContract() bool {
	return a.ContractRegisteredTimestamp != nil
}

func (a *AddressStats) AddToBalance(value *big.Int) {
	a.Balance.Add(a.Balance, value)
	if a.Balance.Cmp(big.NewInt(0)) < 0 {
		log.WithField("address", a.Address).WithField("balance", a.Balance.String()).Warn("Balance is negative")
		a.Balance = big.NewInt(0)
	}
}

type AddressStatsMap struct {
	m            sync.RWMutex
	addressStats map[string]*AddressStats
}

func (asm *AddressStatsMap) AddToBatch(b Batch) {
	for _, stats := range asm.addressStats {
		b.Add(stats, stats.Address, 0)
	}
}

func (asm *AddressStatsMap) GetAddress(s Storage, addr string) *AddressStats {
	addr = strings.ToLower(addr)
	asm.m.Lock()
	defer asm.m.Unlock()
	stats := asm.addressStats[addr]
	if stats != nil {
		return stats
	}

	asm.addressStats[addr] = s.GetAddressStats(addr)

	return asm.addressStats[addr]
}

func (asm *AddressStatsMap) UpdateBalances(s Storage, from, to, valueStr string) {
	value, ok := big.NewInt(0).SetString(valueStr, 0)

	if ok && value.Cmp(big.NewInt(0)) > 0 {
		asm.GetAddress(s, from).AddToBalance(big.NewInt(0).Neg(value))
		asm.GetAddress(s, to).AddToBalance(value)
	}
}

func (asm *AddressStatsMap) UpdateEvents(s Storage, logs []models.EventLog) error {
	if len(logs) > 0 {
		rewardsEvents, err := events.DecodeRewardsTopics(logs)
		if err != nil {
			return err
		}
		for _, event := range rewardsEvents {
			asm.GetAddress(s, common.DposContractAddress).AddToBalance(big.NewInt(0).Neg(event.Value))
			asm.GetAddress(s, event.Account).AddToBalance(event.Value)
		}
	}
	return nil
}

func (asm *AddressStatsMap) GetBalance(addr string) *big.Int {
	addr = strings.ToLower(addr)
	asm.m.RLock()
	defer asm.m.RUnlock()
	stats := asm.addressStats[addr]
	if stats != nil {
		return stats.Balance
	}
	return nil
}

func (asm *AddressStatsMap) AddToBalance(s Storage, addr string, value *big.Int) {
	asm.GetAddress(s, addr).AddToBalance(value)
}

func MakeAddressStatsMap() *AddressStatsMap {
	return &AddressStatsMap{
		addressStats: make(map[string]*AddressStats),
	}
}
