package model

type DataStatistic struct {
	Province     string `gorm:"primaryKey;size:50"`
	IndustryType string
	Quarter      uint64
	OutputValue  float64
}
