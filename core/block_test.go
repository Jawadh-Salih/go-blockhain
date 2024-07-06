package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/Jawadh-Salih/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeaderEncodeDecode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     9872122,
	}

	buf := &bytes.Buffer{}

	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlockEncodeDecode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     9872122,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}

	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}
