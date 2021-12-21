package model

type MapStatistic struct {
	Province     string `gorm:"primaryKey;size:50"`
	IndustryType string
	CountType    string
	Count        string
}
