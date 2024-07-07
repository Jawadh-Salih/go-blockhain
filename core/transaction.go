package core

import (
	"fmt"

	"github.com/Jawadh-Salih/go-blockchain/crypto"
)

// a generic representation of Transctions.
type Transaction struct {
	Data []byte // any kind of arbitrary data.

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

// txn needs to be signed
func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.PublicKey = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("inval transaction signature")
	}

	return nil
}
