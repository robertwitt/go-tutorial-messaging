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

// PersistReadingsQueue name of the queue to persist readings
const PersistReadingsQueue = "PersistReadings"

// WebappSourceExchange name of exchange for webapp sources
const WebappSourceExchange = "WebappSources"

// WebappReadingsExchange name of exchange for webapp readings
const WebappReadingsExchange = "WebappReadings"

// WebappDiscoveryQueue name pf queue for webapp discovery
const WebappDiscoveryQueue = "WebappDiscovery"

// GetChannel get new channel
func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establish connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")

	return conn, ch
}

// GetQueue get new queue
func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(name, false, autoDelete, false, false, nil)
	failOnError(err, "Failed to declare queue")

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
