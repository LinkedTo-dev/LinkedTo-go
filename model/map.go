package model

type Map struct {
	Province     string `gorm:"primaryKey;size:50"`
	PosLatitude  float64
	PosLongitude float64
}
