package main

import (
	cache "kvstore/kv"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	kv := cache.NewDistribuetCache()

	rpc.Register(kv)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err.Error())
	}
	// serves the request by looking
	// to the inner rpc server setting up
	// by Register and HandleHTTP func call
	go http.Serve(listener, nil)
}
