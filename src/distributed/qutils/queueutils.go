package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// SensorDiscoveryExchange name of exchange for sensor discovery
const SensorDiscoveryExchange = "SensorDiscovery"

// SensorListQueue name of the queue for sensors
const SensorListQueue = "SensorList"

// GetChannel get new channel
func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establish connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")

	return conn, ch
}

// GetQueue get new queue
func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(name, false, false, false, false, nil)
	failOnError(err, "Failed to declare queue")

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
