package main

import (
	"fmt"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, _ := zmq.NewContext()
	defer ctx.Term()

	receiver, _ := ctx.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	sender, _ := ctx.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	for {
		msgbytes, _ := receiver.Recv(0)
		fmt.Printf("%s\n", string(msgbytes))

		msec, _ := strconv.ParseInt(string(msgbytes), 10, 64)
		time.Sleep(time.Duration(msec) * time.Millisecond)

		sender.Send("", 0)
	}
}
