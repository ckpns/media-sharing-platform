package services

import (
	"errors"

	"github.com/google/uuid"

	"github.com/ckpns/media-sharing-platform/backend/src/data/models"
	"github.com/ckpns/media-sharing-platform/backend/src/data/repositories"
)

func CreatePost(postToCreate *models.Post) error {

	err := repositories.SavePost(postToCreate)
	return err
}

func DeletePost(postID, userID uuid.UUID) error {

	// check if user is owner
	post, err := repositories.GetPostByID(postID)
	if err != nil {
		return err
	}

	if post.UserID != userID {
		return errors.New("user is not owner of this post")
	}

	// if success, then delete
	err = repositories.DeletePostByID(postID)
	return err
}

func GetPost(postID uuid.UUID) (*models.Post, error) {

	post, err := repositories.GetPostByID(postID)
	if err != nil {
		return &models.Post{}, err
	}

	return post, nil
}

func FavoritePost(postID, userID uuid.UUID) error {

	// check if post exists
	_, err := repositories.GetPostByID(postID)
	if err != nil {
		return err
	}

	// check if user exists
	_, err = repositories.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = repositories.FavoritePost(postID, userID)
	return err
}

func UnfavoritePost(postID, userID uuid.UUID) error {

	// check if post exists
	_, err := repositories.GetPostByID(postID)
	if err != nil {
		return err
	}

	// check if user exists
	_, err = repositories.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = repositories.UnfavoritePost(postID, userID)
	return err
}
