package main

import (
	"examples/chapter_15/rpc_objects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	cal := new(rpc_objects.Args)
	rpc.Register(cal)
	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
