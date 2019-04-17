package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, _ := zmq.NewContext()
	defer ctx.Term()

	receiver, _ := ctx.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind("tcp://*:5558")

	msgbytes, _ := receiver.Recv(0)
	fmt.Println("Received start msg ", string(msgbytes))

	startTime := time.Now().UnixNano()

	for i := 0; i < 100; i++ {
		msgbytes, _ = receiver.Recv(0)
		fmt.Print(".")
	}

	te := time.Now().UnixNano()
	fmt.Printf("Total elapsed time: %d msec\n", (te-startTime)/1e6)
}
