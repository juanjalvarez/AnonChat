package main

type Event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewEvent(t string, d interface{}) *Event {
	return &Event{t, d}
}
