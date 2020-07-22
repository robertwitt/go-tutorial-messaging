package main

import (
	"fmt"

	"github.com/robertwitt/go-tutorial-messaging/src/distributed/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}
