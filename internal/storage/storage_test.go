package storage

import (
	"os"
	"strconv"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func TestGetter(t *testing.T) {
	storage := NewStorage("")
	defer storage.Close()

	addr := MakeEmptyAddressStats("test")
	if err := storage.addToDBTest(addr, addr.Address, 0); err != nil {
		t.Error(err)
	}
	addr1 := MakeEmptyAddressStats("test1")
	if err := storage.addToDBTest(addr1, addr1.Address, 0); err != nil {
		t.Error(err)
	}
	ret := storage.GetAddressStats("test")
	if !ret.isEqual(addr) {
		t.Error("Broken DB")
	}
}

func TestGetObjects(t *testing.T) {
	stor := NewStorage("")
	defer stor.Close()

	sender := "user"
	count := uint64(100)
	for i := uint64(1); i <= count; i++ {
		block := models.Dag{Timestamp: i, Hash: "test" + strconv.FormatUint(i, 10), Level: 0, Sender: sender, TransactionCount: 0}
		if err := stor.addToDBTest(&block, block.Sender, block.Timestamp); err != nil {
			t.Error(err)
		}
	}

	ret, pagination, err := GetObjectsPage[models.Dag](stor, sender, 0, count, count)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, uint64(len(ret)), count)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination, err = GetObjectsPage[models.Dag](stor, sender, 50, 50, 100)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(ret), 50)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(50))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination, err = GetObjectsPage[models.Dag](stor, sender, 0, 25, 100)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(ret), 25)
	assert.True(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(25))
}

func TestStorage(t *testing.T) {
	addr := MakeEmptyAddressStats("test")
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
		if !ret.isEqual(addr) {
			t.Error("Broken DB")
		}
	}
	os.Remove("/tmp/test")
}

func TestCleanStorage(t *testing.T) {
	stats := MakeEmptyAddressStats("test")
	stats.PbftCount = 1
	storage := NewStorage("/tmp/test")
	defer storage.Close()

	if err := storage.addToDBTest(stats, stats.Address, 0); err != nil {
		t.Error(err)
	}

	if err := storage.Clean(); err != nil {
		t.Error(err)
	}

	err := storage.getFromDB(stats, getKey(getPrefix(stats), "test", 0))

	if err == nil {
		t.Error("Clean DB does not work")
		os.Remove("/tmp/test")
	}
}

func TestBatch(t *testing.T) {
	storage := NewStorage("")
	defer storage.Close()

	addr := MakeEmptyAddressStats("test")
	batch := storage.NewBatch()

	batch.AddToBatch(addr, addr.Address, 0)
	addr1 := MakeEmptyAddressStats("test1")
	batch.AddToBatch(addr1, addr1.Address, 0)
	if err := batch.Commit(nil); err != nil {
		t.Error(err)
	}

	ret := storage.GetAddressStats("test")
	if !ret.isEqual(addr) {
		t.Error("Broken DB")
	}

	ret = storage.GetAddressStats("test1")
	if !ret.isEqual(addr1) {
		t.Error("Broken DB")
	}
}
