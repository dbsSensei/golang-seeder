package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"size:255"`
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
