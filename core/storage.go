package core

type Storage interface {
	Put(*Block) error
	// Get()
}

type BlockStorage struct{}

type MemoryStore struct {
}

func (s *MemoryStore) Put(b *Block) error {
	return nil
}
