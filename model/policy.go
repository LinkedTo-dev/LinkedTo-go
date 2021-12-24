package model

type Policy struct {
	ID           uint64 `json:"id"`
	Title        string `json:"title"`
	IndustryType string `json:"industryType"`
	Source       string `json:"source"`
	Organization string `json:"organization"`
	Time         string `json:"time"`
}
