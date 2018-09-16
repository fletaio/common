package store

import (
	"bytes"

	"github.com/tidwall/buntdb"
)

// Bunt TODO
type Bunt struct {
	db *buntdb.DB
}

// NewBunt TODO
func NewBunt(path string) (*Bunt, error) {
	st := new(Bunt)

	db, err := buntdb.Open(path)
	if err != nil {
		return nil, err
	}
	st.db = db

	return st, nil
}

// Close TODO
func (st *Bunt) Close() {
	st.db.Shrink()
	st.db.Close()
}

// Set TODO
func (st *Bunt) Set(key []byte, value []byte) error {
	return st.db.Update(func(tx *buntdb.Tx) error {
		if _, _, err := tx.Set(string(key), string(value), nil); err != nil {
			return err
		}
		return nil
	})
}

// Get TODO
func (st *Bunt) Get(key []byte) ([]byte, error) {
	var data []byte
	if err := st.db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get(string(key))
		if err != nil {
			if err == buntdb.ErrNotFound {
				return ErrNotExistKey
			} else {
				return err
			}
		}
		data = []byte(value)
		return nil
	}); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete TODO
func (st *Bunt) Delete(key []byte) error {
	return st.db.Update(func(tx *buntdb.Tx) error {
		if _, err := tx.Delete(string(key)); err != nil {
			if err == buntdb.ErrNotFound {
				return ErrNotExistKey
			} else {
				return err
			}
		}
		return nil
	})
}

// Scan TODO
func (st *Bunt) Scan(prefix []byte) ([][]byte, [][]byte, error) {
	keys := make([][]byte, 0)
	values := make([][]byte, 0)
	if err := st.db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("", func(k, v string) bool {
			if len(k) >= len(prefix) {
				if bytes.Equal(prefix, []byte(k)[:len(prefix)]) {
					keys = append(keys, []byte(k))
					values = append(values, []byte(v))
				}
			}
			return true
		})
		return nil
	}); err != nil {
		return nil, nil, err
	}
	return keys, values, nil
}
