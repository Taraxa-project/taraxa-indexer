package storage

import (
	"os"
	"strconv"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/models"
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
	ret := MakeEmptyAddressStats("")
	if err := storage.GetFromDB(ret, "test"); err != nil {
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
	for i := 0; i <= count; i++ {
		block := models.Dag{Age: uint64(i), Hash: "test" + strconv.Itoa(i), Level: 0, Sender: sender, TransactionCount: 0}
		if err := stor.AddToDB(&block, block.Sender, block.Age); err != nil {
			t.Error(err)
		}
	}
	ret, err := GetObjectsPage[models.Dag](stor, sender, 0, count)
	if err != nil {
		t.Error(err)
	}
	if len(ret) != count {
		t.Error("Wrong length", len(ret))
	}

	ret, err = GetObjectsPage[models.Dag](stor, sender, 49, 100)
	if err != nil {
		t.Error(err)
	}
	if len(ret) != 50 {
		t.Error("Wrong length", len(ret))
	}
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
		var ret AddressStats
		if err := storage.GetFromDB(&ret, "test"); err != nil {
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

	var ret AddressStats
	if err := storage.GetFromDB(&ret, "test"); err == nil {
		t.Error("Clean DB does not work")
		os.Remove("/tmp/test")
	}
}

func TestBatch(t *testing.T) {
	storage := NewStorage("")
	defer storage.Close()

	addr := MakeEmptyAddressStats("test")
	batch := storage.NewBatch()

	if err := batch.AddToBatch(addr, addr.Address, 0); err != nil {
		t.Error(err)
	}
	addr1 := MakeEmptyAddressStats("test1")
	if err := batch.AddToBatch(addr1, addr1.Address, 0); err != nil {
		t.Error(err)
	}
	if err := batch.Commit(nil); err != nil {
		t.Error(err)
	}

	ret := MakeEmptyAddressStats("")
	if err := storage.GetFromDB(ret, "test"); err != nil {
		t.Error(err)
	}
	if !ret.isEqual(addr) {
		t.Error("Broken DB")
	}

	if err := storage.GetFromDB(ret, "test1"); err != nil {
		t.Error(err)
	}
	if !ret.isEqual(addr1) {
		t.Error("Broken DB")
	}
}
