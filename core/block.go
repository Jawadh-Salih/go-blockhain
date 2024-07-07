package core

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/Jawadh-Salih/go-blockchain/crypto"

	"github.com/Jawadh-Salih/go-blockchain/types"
)

type Header struct {
	Version   uint32
	DataHash  types.Hash // hash of all the transactions
	PrevBlock types.Hash
	Timestamp int64
	Height    uint32

	Nonce uint64
}

// instead of returning a []byte, we pass a generic writer.
func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &h.Nonce)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

// if we want hash the block we can't hash everything.
type Block struct {
	*Header      // we want to keep a references of headers than copies.
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// Cached version of the block hash
	hash types.Hash
}

func NewBlock(h *Header, txns []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txns,
	}
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, dec Encoder[*Block]) error {
	return dec.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	//
	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("invalid block signature")
	}

	return nil
}

func (b *Block) HeaderData() []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	enc.Encode(b.Header)

	return buf.Bytes()
}
