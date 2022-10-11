package repositories

import (
	"github.com/ckpns/media-sharing-platform/backend/src/data"
	"github.com/ckpns/media-sharing-platform/backend/src/data/models"
	"github.com/google/uuid"
)

func SavePost(postToSave *models.Post) error {

	db := data.GetDB()

	err := db.Create(postToSave).Error

	return err
}

func GetPostByID(id uuid.UUID) (*models.Post, error) {

	db := data.GetDB()

	post := models.Post{}

	err := db.Find(&post, id).First(&post).Error
	if err != nil {
		return &models.Post{}, err
	}

	return &post, nil
}

func GetPostsByID(ids []uuid.UUID) ([]models.Post, error) {

	db := data.GetDB()

	posts := []models.Post{}

	err := db.Find(&posts, ids).Error
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}

func DeletePostByID(id uuid.UUID) error {

	db := data.GetDB()

	err := db.Delete(&models.Post{}, id).Error
	return err
}

func FavoritePost(postID, userID uuid.UUID) error {

	db := data.GetDB()

	m := models.FavoritedPost{PostID: postID, UserID: userID}

	err := db.Model(&models.FavoritedPost{}).Create(m).Error
	return err
}

func UnfavoritePost(postID, userID uuid.UUID) error {

	db := data.GetDB()

	m := models.FavoritedPost{PostID: postID, UserID: userID}

	err := db.Model(&models.FavoritedPost{}).Delete(m).Error
	return err
}
