package main

import (
	"net/http"

	"github.com/robertwitt/go-tutorial-messaging/src/distributed/web/controller"
)

func main() {
	controller.Initialize()

	http.ListenAndServe(":3000", nil)
}
