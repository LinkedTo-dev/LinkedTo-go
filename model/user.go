package model

import (
	"time"
)

type User struct {
	Username  string    `gorm:"primaryKey" json:"username"`
	Password  []byte    `gorm:"not null;size:72" json:"password"`
	UserType  string    `json:"userType"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	Tel       string    `json:"tel"`
	SessionID string    `json:"sessionID"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}
