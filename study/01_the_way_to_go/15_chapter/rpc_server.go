package main

import (
	"go_code/study/01_the_way_to_go/rpc_objects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)

	rpc.Register(calc)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Starting RPC-server -listen error: ", err)
	}

	go http.Serve(listener, nil)

	time.Sleep(10000e9)
}
