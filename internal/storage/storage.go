package storage

type BoltStorage struct {
}

func NewBoltStore(path string) (*BoltStorage, error) {
	return &BoltStorage{}, nil
}

func (storage *BoltStorage) Close() {

}
