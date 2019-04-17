package main

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, _ := zmq.NewContext()
	defer ctx.Term()

	sender, _ := ctx.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Bind("tcp://*:5557")

	sink, _ := ctx.NewSocket(zmq.PUSH)
	defer sink.Close()
	sink.Connect("tcp://localhost:5558")

	fmt.Printf("Press Enter when the workers are ready.")

	var line string
	fmt.Scanln(&line)

	fmt.Println("Sending tasks to workers...")

	sink.Send("0", 0)

	rand.Seed(time.Now().UnixNano())

	totalMsec := 0

	for i := 0; i < 100; i++ {
		workload := rand.Intn(100)
		totalMsec += workload
		msg := fmt.Sprintf("%d", workload)
		sender.Send(msg, 0)
	}

	fmt.Printf("Total expected cost: %d msec\n", totalMsec)

	time.Sleep(1e9)
}
