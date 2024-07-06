package types

import (
	"crypto/rand"
	"fmt"
)

type Hash [32]uint8

// hash is to make the refrence for the previous block

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("bytes with length %d should be 32", len(b))
		panic(msg)
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)

	return token
}
func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
