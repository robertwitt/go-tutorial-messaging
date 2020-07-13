package dto

import (
	"encoding/gob"
	"time"
)

// SensorMessage sensor message type
type SensorMessage struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

func init() {
	gob.Register(SensorMessage{})
}
