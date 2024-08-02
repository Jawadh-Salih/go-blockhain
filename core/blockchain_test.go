package core

import (
	"testing"

	"github.com/Jawadh-Salih/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)

	return bc

}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevHeader)
}

func TestBlockchain(t *testing.T) {

	bc := newBlockchainWithGenesis(t)

	// adding the genesis block gives a height
	assert.Equal(t, 0, int(bc.Height()))
	assert.NotNil(t, bc)
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.True(t, bc.HasBlock(0))

}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenblocks := 1000
	for i := 0; i < lenblocks; i++ {
		block := randomBlockWithSig(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
		// assert.Equal(t, uint32(i+1), bc.Height())
	}

	assert.Equal(t, uint32(lenblocks), bc.Height())
	assert.Equal(t, lenblocks+1, len(bc.headers))

	assert.NotNil(t, bc.AddBlock(randomBlock(3, getPrevBlockHash(t, bc, 3))))

}

func TestAddBlockToHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSig(t, 3, types.Hash{})))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlks := 100
	for i := 0; i < lenBlks; i++ {
		block := randomBlockWithSig(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)

	}
}
