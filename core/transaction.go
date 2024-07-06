package core

import "io"

// a generic representation of Transctions.
type Transaction struct {
	Data []byte // any kind of arbitrary data.
}

// txn needs to be signed

func (t *Transaction) EncodeBinary(w io.Writer) error { return nil }
func (t *Transaction) DecodeBinary(r io.Reader) error { return nil }
