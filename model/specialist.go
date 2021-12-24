package model

type Specialist struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	IndustryType string `json:"industryType"`
	Workplace    string `json:"workplace"`
	Job          string `json:"job"`
	Intro        string `json:"intro"`
	Source       string `json:"source"`
}
