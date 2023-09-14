package models

import (
	"time"
)

type Warning struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Water     uint32    `gorm:"not_null" json:"water"`
	Wind      uint32    `gorm:"not_null" json:"wind"`
	Status    string    `gorm:"not_null" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
