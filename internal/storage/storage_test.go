package storage

import (
	"strconv"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func TestAddress(t *testing.T) {
	storage := NewStorage("")
	addr := Address{"test", 0, 0, 0}
	if err := storage.AddToDB(&addr); err != nil {
		t.Error(err)
	}
	addr1 := Address{"test1", 1, 0, 0}
	if err := storage.AddToDB(&addr1); err != nil {
		t.Error(err)
	}
	var ret Address
	if err := storage.GetFromDB(&ret, "test") ; err != nil {
		t.Error(err)
	}
	if ret != addr {
		t.Error("Broken DB")
	}
}

func TestGetObjects(t *testing.T) {
	storage := NewStorage("")
	sender := "user"
	count := 100
	for i := 0; i <= count; i++ {
		block := models.Dag{Age: uint64(i), Hash: "test" + strconv.Itoa(i), Level: 0, Sender: sender, TransactionCount: 0}
		if err := storage.AddToDB(&block); err != nil {
			t.Error(err)
		}
	}
	var block models.Dag
	ret, err := storage.GetObjects(&block, sender, 0 ,count)
	if err != nil {
		t.Error(err)
	}
	if len(ret) != count {
		t.Error("Wrong length", len(ret))
	}

	ret, err = storage.GetObjects(&block, sender, 49, 100)
	if err != nil {
		t.Error(err)
	}
	if len(ret) != 50 {
		t.Error("Wrong length", len(ret))
	}
}
