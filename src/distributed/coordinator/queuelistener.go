package coordinator

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/robertwitt/go-tutorial-messaging/src/distributed/dto"
	"github.com/robertwitt/go-tutorial-messaging/src/distributed/qutils"
	"github.com/streadway/amqp"
)

const url = "amqp://guest:guest@localhost:5672"

// QueueListener queue listener
type QueueListener struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	sources map[string]<-chan amqp.Delivery
	ea      *EventAggregator
}

// NewQueueListener new queue listener
func NewQueueListener() *QueueListener {
	ql := QueueListener{
		sources: make(map[string]<-chan amqp.Delivery),
		ea:      NewEventAggregator(),
	}
	ql.conn, ql.ch = qutils.GetChannel(url)

	return &ql
}

// DiscoverSensors discover sensors
func (ql *QueueListener) DiscoverSensors() {
	ql.ch.ExchangeDeclare(qutils.SensorDiscoveryExchange, "fanout", false, false, false, false, nil)

	ql.ch.Publish(qutils.SensorDiscoveryExchange, "", false, false, amqp.Publishing{})
}

// ListenForNewSource listen for new source
func (ql *QueueListener) ListenForNewSource() {
	q := qutils.GetQueue("", ql.ch)
	ql.ch.QueueBind(q.Name, "", "amq.fanout", false, nil)

	msgs, _ := ql.ch.Consume(q.Name, "", true, false, false, false, nil)

	ql.DiscoverSensors()

	for msg := range msgs {
		sourceChan, _ := ql.ch.Consume(string(msg.Body), "", true, false, false, false, nil)

		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan
			go ql.AddListener(sourceChan)
		}
	}
}

// AddListener add listener
func (ql *QueueListener) AddListener(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(dto.SensorMessage)
		d.Decode(sd)

		fmt.Printf("Received message: %v\n", sd)

		ed := EventData{
			Name:      sd.Name,
			Timestamp: sd.Timestamp,
			Value:     sd.Value,
		}

		ql.ea.PublishEvent("MessageReceived_"+msg.RoutingKey, ed)
	}
}