package model

import "time"

type Policy struct {
	ID           uint64
	Title        string
	Organization string
	Time         time.Time
}
