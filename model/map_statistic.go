package model

type MapStatistic struct {
	Province     string `gorm:"primaryKey;size:50" json:"province"`
	IndustryType string `json:"industryType"`
	CountType    string `json:"countType"`
	Count        string `json:"count"`
}
