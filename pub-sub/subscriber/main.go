package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, _ := zmq.NewContext()
	socket, _ := ctx.NewSocket(zmq.SUB)
	defer ctx.Term()
	defer socket.Close()

	totalTemp := 0
	filter := "59937"

	if len(os.Args) > 1 {
		filter = string(os.Args[1])
	}

	fmt.Printf("Collecting updates from wather server for %s...\n", filter)
	socket.SetSubscribe(filter)
	socket.Connect("tcp://localhost:5566")

	for i := 0; i < 101; i++ {
		datapt, _ := socket.Recv(0)
		temps := strings.Split(datapt, " ")
		temp, err := strconv.ParseInt(temps[1], 10, 64)
		if nil == err {
			totalTemp += int(temp)
		}
	}

	fmt.Printf("average temperature for zipcode %s was %dF \n\n", filter, totalTemp/100)

}
