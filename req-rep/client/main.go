package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

const (
	endpoint = "tcp://localhost:5555"
)

func main() {
	ctx, _ := zmq.NewContext()
	socket, _ := ctx.NewSocket(zmq.REQ)
	defer ctx.Term()
	defer socket.Close()

	fmt.Printf("connecting to hello world server...")
	socket.Connect("tcp://localhost:5555")

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Hello %d", i)
		socket.Send(msg, 0)

		reply, _ := socket.Recv(0)
		println("Received: ", string(reply))
	}
}
