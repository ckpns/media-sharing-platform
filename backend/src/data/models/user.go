package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"` // Warning: gen_random_uuid() only works with postgreSQL 14+
	Username       string          `gorm:"unique;not null;"`
	Password       string          `gorm:"not null;"`
	Posts          []Post          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	FavoritedPosts []FavoritedPost `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time       `gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime"`
}
