package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"gogo-server.com/dcx/codec"
	"gogo-server.com/dcx/server"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	server.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple geerpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	err := json.NewEncoder(conn).Encode(server.DefaultOption)

	if err != nil {
		log.Panicln("Encode DefaultOption:", err)
	}
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		err := cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		if err != nil {
			log.Println("Write:", err)
		}
		_ = cc.ReadHeader(h)
		var reply string
		err = cc.ReadBody(&reply)
		if err != nil {
			log.Println("ReadBody:", err)
		}
		log.Println("reply:", reply)
	}
	time.Sleep(time.Second * 10)
}
