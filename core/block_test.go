package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/Jawadh-Salih/go-blockchain/crypto"
	"github.com/Jawadh-Salih/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {
	b := randomBlock(0)
	privKey := crypto.GeneratePrivateKey()
	fmt.Println(b.Hash(BlockHasher{}))

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

	// assert.False(t, h.IsZero())
}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(0)
	privKey := crypto.GeneratePrivateKey()

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	// if the PublicKey is altered.
	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

// random block of given height
func randomBlock(height uint32) *Block {
	header := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     9872122,
	}

	tx := Transaction{
		Data: []byte("Test"),
	}

	return NewBlock(header, []Transaction{tx})
}
