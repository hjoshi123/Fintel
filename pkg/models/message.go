package models

type Message struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}
