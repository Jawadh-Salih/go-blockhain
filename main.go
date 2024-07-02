package main

import (
	"time"

	"github.com/Jawadh-Salih/go-blockchain/network"
)

// server
// transport layer, tcp, udp
// block
// tx
// key pairs

func main() {

	localTr := network.NewLocalTransport("Local")
	remoteTr := network.NewLocalTransport("Remote")

	localTr.Connect(remoteTr)
	remoteTr.Connect(localTr)

	go func() {
		for {
			remoteTr.SendMessage(localTr.Addr(), []byte(`Hello local`))
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOps{
		Transports: []network.Transport{localTr, remoteTr},
	}

	s := network.NewServer(opts)

	s.Start()
}
