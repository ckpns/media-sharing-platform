package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"` // Warning: gen_random_uuid() only works with postgreSQL 14+
	Description string          `gorm:"required;not null;"`
	MediaPath   string          `gorm:"required;not null;"`
	FavoritedBy []FavoritedPost `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`
	UserID      uuid.UUID       `gorm:"type:uuid;not null;"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
}
