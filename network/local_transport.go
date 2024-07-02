package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan Rpc
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport // to track the peers
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan Rpc, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan Rpc {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s couldn't send a message to %s", t.addr, to)
	}

	peer.consumeCh <- Rpc{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}
