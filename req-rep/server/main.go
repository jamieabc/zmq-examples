package main

import (
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	endpoint = "tcp://*:5555"
)

func main() {
	ctx, _ := zmq.NewContext()
	socket, _ := ctx.NewSocket(zmq.REP)
	defer ctx.Term()
	defer socket.Close()
	socket.Bind(endpoint)

	for {
		msg, _ := socket.Recv(0)
		println("Received ", string(msg))

		time.Sleep(time.Second)

		reply := "World"
		socket.Send(reply, 0)
	}
}
