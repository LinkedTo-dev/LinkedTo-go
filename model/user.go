package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	Username  string `gorm:"primaryKey"`
	Password  []byte `gorm:"not null;size:72"`
	UserType  string
	Gender    string
	Email     string
	Tel       string
	SessionID string
	SessionId uuid.UUID `gorm:"not null;uniqueIndex;size:36"`
	UpdatedAt time.Time `gorm:"not null"`
}
