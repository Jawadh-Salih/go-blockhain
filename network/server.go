package network

import (
	"fmt"
	"time"
)

type ServerOps struct {
	Transports []Transport
}

// container that includes every modules
type Server struct {
	ServerOps

	rpcCh  chan Rpc
	quitCh chan struct{}
}

func NewServer(opts ServerOps) *Server {
	return &Server{
		ServerOps: opts,
		rpcCh:     make(chan Rpc),
		quitCh:    make(chan struct{}, 1),
	}

}

func (s *Server) Start() {
	s.initTranports()

	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("I am checking in")
		}
	}

	fmt.Println("Server shutdown")
}

func (s *Server) initTranports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			// keep consuming message channels forever
			for rpc := range tr.Consume() {
				// handle the message
				// the message that comes here
				// are not thread safe
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
