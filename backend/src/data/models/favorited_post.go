package models

import (
	"github.com/google/uuid"
)

type FavoritedPost struct {
	PostID uuid.UUID `gorm:"type:uuid;column:post_id;primary_key;auto_increment:false;"`
	UserID uuid.UUID `gorm:"type:uuid;column:user_id;primary_key;auto_increment:false;"`
}
