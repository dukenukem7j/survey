package model

import (
	"time"
)

// Employee is fucking understandable
type Employee struct {
	ID        uint
	FirstName string
	LastName  string
	StartDate time.Time
	Status    string
}

type TimeOff struct {
	Type      string
	Amount    float32
	StartDate time.Time
	Status    string
}
