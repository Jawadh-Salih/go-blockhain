package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypairSignAndVerifyValid(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	msg := []byte("Test message")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeypairSignAndVerifyInValid(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("Test message")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrvKey := GeneratePrivateKey()
	otherPubKey := otherPrvKey.PublicKey()
	othermsg := []byte("Test message 2")

	assert.False(t, sig.Verify(otherPubKey, msg))
	assert.False(t, sig.Verify(pubKey, othermsg))
}
