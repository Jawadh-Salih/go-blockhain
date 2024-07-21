package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {

	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already has block (%d) with hash %s", b.Height, b.Hash(BlockHasher{}))
	}

	// if err := b.Verify(); err != nil {
	// 	return err
	// }
	// check  the signature whether it's valid. (but expensive compute)
	return nil
}
