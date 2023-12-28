package main

import (
	"log"
	"net"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	

}
func main() {
	log.SetFlags(0)
	addr := make(chan string)
	log.Println(addr)
}
