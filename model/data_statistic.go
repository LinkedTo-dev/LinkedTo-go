package model

type DataStatistic struct {
	Province     string  `json:"province"`
	IndustryType string  `json:"industryType"`
	Quarter      string  `json:"quarter"`
	OutputValue  float64 `json:"outputValue"`
}
