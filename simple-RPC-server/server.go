package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"

	shared "github.com/zquangu112z/simple-RPC/simple-RPC-shared-struct"
)

type MyArith struct{}

func (t MyArith) Multiply(args shared.Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println("Executing Multiply")
	return nil
}

func (t MyArith) Divide(args shared.Args, quo *shared.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	fmt.Println("Executing Devide")
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func registerArith(server *rpc.Server, arith shared.Arith) {
	// registers Arith interface by name of `Arithmetic`.
	// If you want this name to be same as the type name, you
	// can use server.Register instead.
	server.RegisterName("Arithmetic", arith)
}

func main() {

	//Creating an instance of struct which implement Arith interface
	var arith MyArith

	// Register a new rpc server (In most cases, you will use default server only)
	// And register struct we created above by name "Arith"
	// The wrapper method here ensures that only structs which implement Arith interface
	// are allowed to register themselves.
	server := rpc.NewServer()
	registerArith(server, arith)

	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// This statement links rpc server to the socket, and allows rpc server to accept
	// rpc request coming from that socket.
	server.Accept(l)
}
