package main

import (
	"log"
	"net"

	pkg "github.com/arsmuradyan/styxedge/pkg"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	addresses := []string{
		"127.0.0.1:3000",
		"127.0.0.1:3001",
		"127.0.0.1:3002",
	}

	serverPool := pkg.ServerPool{}
	for _, address := range addresses {
		addr, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			panic(err)
		}
		serverPool.AddBackend(addr)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go pkg.Proxy(conn, serverPool.GetNextPeer().Address())
	}
}
