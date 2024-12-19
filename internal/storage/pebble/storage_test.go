package pebble

import (
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/transaction"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/nleeper/goment"
	"github.com/stretchr/testify/assert"
)

func TestGetter(t *testing.T) {
	st := NewStorage("")
	defer st.Close()

	addr := storage.MakeEmptyAddressStats("test")
	if err := st.addToDBTest(addr, addr.Address, 0); err != nil {
		t.Error(err)
	}
	addr1 := storage.MakeEmptyAddressStats("test1")
	if err := st.addToDBTest(addr1, addr1.Address, 0); err != nil {
		t.Error(err)
	}
	ret := st.GetAddressStats("test")
	if !ret.IsEqual(addr) {
		t.Error("Broken DB")
	}
}

func TestGetObjects(t *testing.T) {
	st := NewStorage("")
	defer st.Close()

	sender := "user"
	count := uint64(100)

	stats := storage.MakeEmptyAddressStats(sender)
	stats.DagsCount = count
	if err := st.addToDBTest(stats, stats.Address, 0); err != nil {
		t.Error(err)
	}

	for i := uint64(1); i <= count; i++ {
		block := models.Dag{Timestamp: i, Hash: "test" + strconv.FormatUint(i, 10), Level: 0, TransactionCount: 0}
		if err := st.addToDBTest(&block, sender, block.Timestamp); err != nil {
			t.Error(err)
		}
	}
	ret, pagination := storage.GetObjectsPage[models.Dag](st, sender, 0, uint64(count))
	assert.Equal(t, count, uint64(len(ret)))
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = storage.GetObjectsPage[models.Dag](st, sender, 50, 100)
	assert.Equal(t, len(ret), 50)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(50))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = storage.GetObjectsPage[models.Dag](st, sender, 0, 25)
	assert.Equal(t, len(ret), 25)
	assert.True(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(25))
}

func TestGetPaginatedWeekStats(t *testing.T) {
	stor := NewStorage("")
	defer stor.Close()

	tn, _ := goment.New()
	weekStats := stor.GetWeekStats(int32(tn.ISOWeekYear()), int32(tn.ISOWeek()))

	count := uint64(100)
	for i := uint64(1); i <= count; i++ {
		block := models.Pbft{Author: "user" + strconv.FormatUint(i, 10), Hash: "test" + strconv.FormatUint(i, 10), Number: i, Timestamp: uint64(tn.ToUnix()), TransactionCount: 0}
		weekStats.AddPbftBlock(&block)
	}

	ret, pagination := weekStats.GetPaginated(0, count)
	assert.Equal(t, uint64(len(ret)), count)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = weekStats.GetPaginated(50, 50)
	assert.Equal(t, len(ret), 50)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(50))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = weekStats.GetPaginated(0, 25)
	assert.Equal(t, len(ret), 25)
	assert.True(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(25))
}

func TestStorage(t *testing.T) {
	addr := storage.MakeEmptyAddressStats("test")
	{
		storage := NewStorage("/tmp/test")
		if err := storage.addToDBTest(addr, addr.Address, 0); err != nil {
			t.Error(err)
		}
		storage.Close()
	}
	{
		storage := NewStorage("/tmp/test")
		defer storage.Close()
		ret := storage.GetAddressStats("test")
		if !ret.IsEqual(addr) {
			t.Error("Broken DB")
		}
	}
	os.Remove("/tmp/test")
}

func TestCleanStorage(t *testing.T) {
	stats := storage.MakeEmptyAddressStats("test")
	stats.PbftCount = 1
	st := NewStorage("/tmp/test")
	defer st.Close()

	if err := st.addToDBTest(stats, stats.Address, 0); err != nil {
		t.Error(err)
	}

	if err := st.Clean(); err != nil {
		t.Error(err)
	}

	err := st.GetFromDB(stats, getKey(GetPrefix(stats), "test", 0))

	if err == nil {
		t.Error("Clean DB does not work")
		os.Remove("/tmp/test")
	}
}

func TestBatch(t *testing.T) {
	st := NewStorage("")
	defer st.Close()

	addr := storage.MakeEmptyAddressStats("test")
	batch := st.NewBatch()

	batch.Add(addr, addr.Address, 0)
	addr1 := storage.MakeEmptyAddressStats("test1")
	batch.Add(addr1, addr1.Address, 0)
	batch.CommitBatch()

	ret := st.GetAddressStats("test")
	if !ret.IsEqual(addr) {
		t.Error("Broken DB")
	}

	ret = st.GetAddressStats("test1")
	if !ret.IsEqual(addr1) {
		t.Error("Broken DB")
	}
}

func TestAccountsBatch(t *testing.T) {
	st := NewStorage("")
	defer st.Close()

	accounts := []storage.Account{
		{
			Address: "0x1111111111111111111111111111111111111111",
			Balance: big.NewInt(100),
		},
		{
			Address: "0x0DC0d841F962759DA25547c686fa440cF6C28C61",
			Balance: big.NewInt(50),
		},
	}

	batch := st.NewBatch()
	batch.SaveAccounts(accounts)
	batch.CommitBatch()

	ret := st.GetAccounts()

	if len(ret) != len(accounts) {
		t.Error("Broken DB")
	}

}

func TestTxByHash(t *testing.T) {
	st := NewStorage("")
	defer st.Close()

	tx := storage.Transaction{
		Hash:  "0x111111",
		From:  "0x222222",
		Value: big.NewInt(100),
		To:    "0x00000000000000000000000000000000000000fe",
		Input: "0x5c19a95c000000000000000000000000ed4d5f4f3641cbc056e466d15dbe2403e38056f8",
	}

	batch := st.NewBatch()
	batch.AddSingleKey(tx, tx.Hash)
	batch.CommitBatch()

	ret := st.GetTransactionByHash(tx.Hash)
	err := transaction.DecodeTransaction(&ret)

	assert.NoError(t, err)
	assert.Equal(t, tx.Hash, ret.Hash)
	assert.Equal(t, tx.Input, ret.Input)
	assert.Equal(t, []any{"0xed4d5f4f3641cbc056e466d15dbe2403e38056f8"}, ret.Calldata.Params)
}
