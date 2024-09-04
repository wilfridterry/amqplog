package domain

import "time"

type MessageLog struct {
	Entity    string    `json:"entity"`
	Action    string    `json:"action"`
	EntityID  int64     `json:"entity_id"`
	Timestamp time.Time `json:"timestamp"`
}