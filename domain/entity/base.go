package entity

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Base struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
