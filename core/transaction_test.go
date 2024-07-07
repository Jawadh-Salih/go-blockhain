package core

import (
	"testing"

	"github.com/Jawadh-Salih/go-blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransachtion(t *testing.T) {
	data := []byte("Test message")
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)

}

func TestVerifyTransaction(t *testing.T) {
	data := []byte("Test message")
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}
