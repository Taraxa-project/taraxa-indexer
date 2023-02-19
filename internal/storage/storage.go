package storage

import (
	"io"
	"log"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
)

type Storage struct {
	DB *pebble.DB
}

func NewStorage(file string) *Storage {
	db, err := pebble.Open(file, &pebble.Options{FS: vfs.NewMem()})
	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		DB: db,
	}
}

func (s *Storage) Add(key, value []byte) error {
	err := s.DB.Set(key, value, pebble.Sync)
	return err
}

func (s *Storage) Get(key []byte) ([]byte, io.Closer, error) {
	value, closer, err := s.DB.Get(key)
	if err != nil {
		return nil, nil, err
	}

	return value, closer, nil
}

func (s *Storage) Find(prefix []byte) *pebble.Iterator {
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil // no upper-bound
	}

	prefixIterOptions := func(prefix []byte) *pebble.IterOptions {
		return &pebble.IterOptions{
			LowerBound: prefix,
			UpperBound: keyUpperBound(prefix),
		}
	}

	iter := s.DB.NewIter(prefixIterOptions(prefix))
	return iter
}
