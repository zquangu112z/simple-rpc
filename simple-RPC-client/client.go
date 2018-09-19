package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	shared "github.com/zquangu112z/simple-rpc/simple-RPC-shared-struct"
)

func main() {
	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	conn, err := net.Dial("tcp", ":1337")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}

	client := rpc.NewClient(conn)

	args := shared.Args{5, 6}

	var reply_quo shared.Quotient
	client.Call("Arithmetic.Divide", args, &reply_quo)
	fmt.Println(reply_quo)

	var reply_mul int
	client.Call("Arithmetic.Multiply", args, &reply_mul)
	fmt.Println(reply_mul)

	var reply_greet string
	client.Call("Arithmetic.Greet", "Nicholas", &reply_greet)
	fmt.Println(reply_greet)
}
