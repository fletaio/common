package store

// Store TODO
type Store interface {
	Close()
	Set(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Scan(prefix []byte) ([][]byte, [][]byte, error)
}
