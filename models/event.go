package models

import (
	"time"
)

type Event struct {
	Action            string
	ActionDescription string
	Name              string
	Timestamp         time.Time
}