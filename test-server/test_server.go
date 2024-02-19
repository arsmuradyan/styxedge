package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	str string
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, h.str)
}
func test(c string) http.Handler {
	return Handler{
		str: c,
	}
}

func main() {

	server1 := &http.Server{Addr: ":3000", Handler: test("First Server")}
	server2 := &http.Server{Addr: ":3001", Handler: test("Second Server")}
	server3 := &http.Server{Addr: ":3002", Handler: test("Third Server")}

	// to tes t graceful termination of tcp sessions
	server1.SetKeepAlivesEnabled(false)
	server2.SetKeepAlivesEnabled(false)
	server3.SetKeepAlivesEnabled(false)

	go server1.ListenAndServe()
	go server2.ListenAndServe()
	server3.ListenAndServe()
}
