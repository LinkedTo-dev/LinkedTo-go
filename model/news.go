package model

import "time"

type News struct {
	ID           uint64
	IndustryType string
	Title        string
	Source       string
	Time         time.Time
}
