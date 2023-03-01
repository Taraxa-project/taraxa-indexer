package storage

import (
	"fmt"
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
	if err := storage.AddToDB(addr, addr.Address, 0); err != nil {
		t.Error(err)
	}
	addr1 := MakeEmptyAddressStats("test1")
	if err := storage.AddToDB(addr1, addr1.Address, 0); err != nil {
		t.Error(err)
	}
	ret, err := storage.GetAddressStats("test")
	if err != nil {
		t.Error(err)
	}
	if !ret.isEqual(addr) {
		t.Error("Broken DB")
	}
}

func TestGetObjects(t *testing.T) {
	stor := NewStorage("")
	defer stor.Close()

	sender := "user"
	count := 100
	for i := 1; i <= count; i++ {
		block := models.Dag{Timestamp: uint64(i), Hash: "test" + strconv.Itoa(i), Level: 0, Sender: sender, TransactionCount: 0}
		if err := stor.AddToDB(&block, block.Sender, block.Timestamp); err != nil {
			t.Error(err)
		}
	}
	ret, pagination, err := GetObjectsPage[models.Dag](stor, sender, 0, uint64(count))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(ret), count)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(100))
	assert.Equal(t, pagination.End, uint64(1))

	ret, pagination, err = GetObjectsPage[models.Dag](stor, sender, 50, 100)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(ret), 50)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(50))
	assert.Equal(t, pagination.End, uint64(1))
}

func TestStorage(t *testing.T) {
	addr := MakeEmptyAddressStats("test")
	{
		storage := NewStorage("/tmp/test")
		if err := storage.AddToDB(addr, addr.Address, 0); err != nil {
			t.Error(err)
		}
		storage.Close()
	}
	{
		storage := NewStorage("/tmp/test")
		defer storage.Close()
		ret, err := storage.GetAddressStats("test")
		if err != nil {
			t.Error(err)
		}
		if !ret.isEqual(addr) {
			t.Error("Broken DB")
		}
	}
	os.Remove("/tmp/test")
}

func TestCleanStorage(t *testing.T) {
	addr := MakeEmptyAddressStats("test")
	storage := NewStorage("/tmp/test")
	defer storage.Close()

	if err := storage.AddToDB(addr, addr.Address, 0); err != nil {
		t.Error(err)
	}

	if err := storage.Clean(); err != nil {
		t.Error(err)
	}

	_, err := storage.GetAddressStats("test")
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

	ret, err := storage.GetAddressStats("test")
	if err != nil {
		t.Error(err)
	}
	if !ret.isEqual(addr) {
		t.Error("Broken DB")
	}

	ret, err = storage.GetAddressStats("test1")
	if err != nil {
		t.Error(err)
	}
	if !ret.isEqual(addr1) {
		t.Error("Broken DB")
	}
}

func TestParseKeyIndex(t *testing.T) {
	v := uint64(28)
	key := "test0000000" + fmt.Sprint(v)
	prefix := "test"
	res := ParseKeyIndex(key, prefix)
	if v != res {
		t.Error("ParseKeyIndex ", v, "!=", res)
	}
}
