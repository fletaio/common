package store

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// LevelDB TODO
type LevelDB struct {
	db *leveldb.DB
}

// NewLevelDB TODO
func NewLevelDB(path string) (*LevelDB, error) {
	st := new(LevelDB)

	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	st.db = db

	return st, nil
}

// Close TODO
func (st *LevelDB) Close() {
	st.db.Close()
}

// Set TODO
func (st *LevelDB) Set(key []byte, value []byte) error {
	return st.db.Put(key, value, nil)
}

// Get TODO
func (st *LevelDB) Get(key []byte) ([]byte, error) {
	if data, err := st.db.Get(key, nil); err != nil {
		if err == leveldb.ErrNotFound {
			return nil, ErrNotExistKey
		} else {
			return nil, err
		}
	} else {
		return data, nil
	}
}

// Delete TODO
func (st *LevelDB) Delete(key []byte) error {
	return st.db.Delete(key, nil)
}

// Scan TODO
func (st *LevelDB) Scan(prefix []byte) ([][]byte, [][]byte, error) {
	keys := make([][]byte, 0)
	values := make([][]byte, 0)
	var r *util.Range
	if prefix != nil {
		r = util.BytesPrefix(prefix)
	}
	iter := st.db.NewIterator(r, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		keys = append(keys, key)
		values = append(values, value)
	}
	iter.Release()
	return keys, values, iter.Error()
}
