package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Client struct model
type Client struct {
	gorm.Model
	ID        uuid.UUID `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_At"`
}
