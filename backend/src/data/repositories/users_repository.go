package repositories

import (
	"github.com/ckpns/media-sharing-platform/backend/src/data"
	"github.com/ckpns/media-sharing-platform/backend/src/data/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SaveUser(userToSave *models.User) error {

	db := data.GetDB()

	err := db.Save(userToSave).Error
	return err
}

func DeleteUserByID(id uuid.UUID) error {

	db := data.GetDB()

	err := db.Delete(&models.User{}, id).Error
	return err
}

func GetUserByID(id uuid.UUID) (*models.User, error) {

	db := data.GetDB()

	user := models.User{}

	err := db.Find(&models.User{}, id).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func GetUserByIdWithPostsPaginated(id uuid.UUID, resultsPerPage, pageNum uint) (*models.User, error) {

	db := data.GetDB()

	user := models.User{}
	offset := (pageNum - 1) * resultsPerPage

	err := db.Preload("Posts", func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(int(offset)).Limit(int(resultsPerPage))
	}).Find(&user, id).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func GetUserByIdWithFavoritesPaginated(id uuid.UUID, resultsPerPage, pageNum uint) (*models.User, error) {

	db := data.GetDB()

	user := models.User{}
	offset := (pageNum - 1) * resultsPerPage

	err := db.Preload("FavoritedPosts", func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(int(offset)).Limit(int(resultsPerPage))
	}).Find(&user, id).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {

	db := data.GetDB()

	user := models.User{}

	err := db.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}
