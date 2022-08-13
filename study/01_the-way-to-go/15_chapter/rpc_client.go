package main

import (
	"fmt"
	"go_code/study/01_the-way-to-go/rpc_objects"
	"log"
	"net/rpc"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("Error dialing: ", err)
	}

	args := &rpc_objects.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)

	if err != nil {
		log.Fatal("Args error: ", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
