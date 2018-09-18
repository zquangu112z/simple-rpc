package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	shared "github.com/zquangu112z/simple-RPC/simple-RPC-shared-struct"
)

type ArithClient struct {
	client *rpc.Client
}

func (t ArithClient) divide(a, b int) shared.Quotient {
	args := &shared.Args{a, b}
	var reply shared.Quotient
	err := t.client.Call("Arithmetic.Divide", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func (t ArithClient) multiply(a, b int) int {
	args := &shared.Args{a, b}
	var reply int
	err := t.client.Call("Arithmetic.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func main() {

	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	conn, err := net.Dial("tcp", ":1337")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}

	// Create a struct, that mimics all methods provided by interface.
	// It is not compulsory, we are doing it here, just to simulate a traditional method call.
	arith := ArithClient{client: rpc.NewClient(conn)}

	fmt.Println(arith.multiply(5, 6))
	fmt.Println(arith.divide(500, 10))
}
