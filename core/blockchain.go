package core

import "fmt"

type Blockchain struct {
	store     Storage
	headers   []*Header // keep everything into memory RAM.
	validator Validator
}

// genesis block is the first block in the chain.
// it's a predefined block
// bc is basically a state machine

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   &MemoryStore{},
	}

	// add the genesis block
	if err := bc.addBlockWithoutValidation(genesis); err != nil {
		return nil, err
	}

	bc.validator = NewBlockValidator(bc) // make the new validator.

	return bc, nil
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

// [0 1 2 3] => 4 headers. But height is 3
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1) // why - 1
}

// need a validator to valdate a block
func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	if err := bc.addBlockWithoutValidation(b); err != nil {
		return err
	}

	return nil
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return bc.Height() >= height
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	return bc.store.Put(b)
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given Height (%d) is too high", height)
	}

	return bc.headers[height], nil
}
