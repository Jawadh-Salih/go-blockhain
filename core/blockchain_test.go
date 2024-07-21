package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(9))
	assert.Nil(t, err)

	return bc

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

func TestAddBloc(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenblocks := 1000
	for i := 0; i < lenblocks; i++ {
		block := randomBlockWithSig(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(block))
		// assert.Equal(t, uint32(i+1), bc.Height())
	}

	assert.Equal(t, uint32(lenblocks), bc.Height())
	assert.Equal(t, lenblocks+1, len(bc.headers))

	assert.NotNil(t, bc.AddBlock(randomBlock(3)))

}
