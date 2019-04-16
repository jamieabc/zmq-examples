package main

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, _ := zmq.NewContext()
	socket, _ := zmq.NewSocket(zmq.PUB)
	defer ctx.Term()
	defer socket.Close()

	socket.Bind("tcp://*:5566")
	socket.Bind("ipc://weather.ipc")

	rand.Seed(time.Now().UnixNano())

	for {
		zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		msg := fmt.Sprintf("%d %d %d", zipcode, temperature, relhumidity)

		socket.Send(msg, 0)
	}
}
