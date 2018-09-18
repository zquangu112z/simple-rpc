package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	shared "github.com/zquangu112z/simple-RPC/simple-RPC-shared-struct"
)

func main() {
	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	conn, err := net.Dial("tcp", ":1337")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}

	args := shared.Args{5, 6}
	var reply_mul int
	var reply_quo shared.Quotient
	client := rpc.NewClient(conn)
	client.Call("Arithmetic.Divide", args, &reply_quo)
	fmt.Println(reply_quo)
	client.Call("Arithmetic.Multiply", args, &reply_mul)
	fmt.Println(reply_mul)
}
