package model

type News struct {
	ID           uint64 `json:"id"`
	IndustryType string `json:"industryType"`
	Title        string `json:"title"`
	Source       string `json:"source"`
	Time         string `json:"time"`
}
