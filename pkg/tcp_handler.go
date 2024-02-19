package pkg

import (
	"io"
	"log"
	"net"
	"sync"
)

func Proxy(conn net.Conn, targetAddr net.Addr) {

	target, err := net.Dial("tcp", targetAddr.String())
	if err != nil {
		log.Fatal("unable to contact to target")
	}
	defer target.Close()
	defer conn.Close()
	var wgt sync.WaitGroup

	// Create TeeReader to log payloads

	wgt.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		_, err := io.Copy(target, conn)
		if err != nil {
			log.Println("unable to forwared to target ", err)
		}
	}(&wgt)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		_, err = io.Copy(conn, target)
		if err != nil {
			log.Println("unable to forwared to client ", err)
		}
	}(&wgt)
	wgt.Wait()
}
