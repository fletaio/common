package store

import (
	"os"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger"
)

// Badger TODO
type Badger struct {
	db       *badger.DB
	lockfile *os.File
	ticker   *time.Ticker
}

// NewBadger TODO
func NewBadger(path string) (*Badger, error) {
	st := new(Badger)

	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	opts.Truncate = true
	lockfilePath := filepath.Join(opts.Dir, "LOCK")

	os.Remove(lockfilePath)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	st.db = db

	lockfile, err := os.OpenFile(lockfilePath, os.O_EXCL, 0)
	if err != nil {
		return nil, err
	}
	st.lockfile = lockfile

	{
	again:
		if err := db.RunValueLogGC(0.7); err != nil {
		} else {
			goto again
		}
	}

	st.ticker = time.NewTicker(5 * time.Minute)
	go func() {
		for range st.ticker.C {
		again:
			if err := db.RunValueLogGC(0.7); err != nil {
			} else {
				goto again
			}
		}
	}()

	return st, nil
}

// Close TODO
func (st *Badger) Close() {
	st.ticker.Stop()
	st.lockfile.Close()
	st.db.Close()
}

// Set TODO
func (st *Badger) Set(key []byte, value []byte) error {
	return st.db.Update(func(txn *badger.Txn) error {
		if err := txn.Set(key, value); err != nil {
			return err
		}
		return nil
	})
}

// Get TODO
func (st *Badger) Get(key []byte) ([]byte, error) {
	var data []byte
	if err := st.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return ErrNotExistKey
			} else {
				return err
			}
		}
		value, err := item.Value()
		if err != nil {
			return err
		}
		data = value
		return nil
	}); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete TODO
func (st *Badger) Delete(key []byte) error {
	return st.db.Update(func(txn *badger.Txn) error {
		if err := txn.Delete(key); err != nil {
			if err == badger.ErrKeyNotFound {
				return ErrNotExistKey
			} else {
				return err
			}
		}
		return nil
	})
}

// Scan TODO
func (st *Badger) Scan(prefix []byte) ([][]byte, [][]byte, error) {
	keys := make([][]byte, 0)
	values := make([][]byte, 0)
	if err := st.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		if len(prefix) > 0 {
			prefix := []byte("1234")
			for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
				keys = append(keys, it.Item().Key())
				if value, err := it.Item().Value(); err != nil {
					return err
				} else {
					values = append(values, value)
				}
			}
		} else {
			for it.Rewind(); it.Valid(); it.Next() {
				keys = append(keys, it.Item().Key())
				if value, err := it.Item().Value(); err != nil {
					return err
				} else {
					values = append(values, value)
				}
			}
		}
		return nil
	}); err != nil {
		return nil, nil, err
	}
	return keys, values, nil
}
