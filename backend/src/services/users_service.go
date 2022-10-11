package services

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/ckpns/media-sharing-platform/backend/src/common"
	"github.com/ckpns/media-sharing-platform/backend/src/data/models"
	"github.com/ckpns/media-sharing-platform/backend/src/data/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(userToCreate *models.User) error {

	err := encryptPassword(userToCreate)
	if err != nil {
		return err
	}

	err = repositories.SaveUser(userToCreate)
	return err
}

func encryptPassword(user *models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func LoginUser(username, password string) (string, error) {

	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = checkPassword(user, password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func checkPassword(user *models.User, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func generateToken(id uuid.UUID) (string, error) {

	config := common.GetConfig()
	expirationInMinutes := config.AUTH.ExpirationMinutes
	hmacSecret := config.AUTH.HmacSecret

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expirationInMinutes)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(hmacSecret))
}

func DeleteUser(id uuid.UUID) error {

	// Warning: Posts associated with this user get automatically deleted (CASCADE)
	// The files though don't get deleted.
	// Maybe create a file service? But what if the file service isn't available and the record is already deleted from DB?
	// TODO: Research message brokers & microservice architecture

	_, err := repositories.GetUserByID(id)
	if err != nil {
		return err
	}

	err = repositories.DeleteUserByID(id)
	return err
}

func GetUserPostsPaginated(userID uuid.UUID, resultsPerPage, pageNum uint) ([]models.Post, error) {

	user, err := repositories.GetUserByIdWithPostsPaginated(userID, resultsPerPage, pageNum)
	if err != nil {
		return []models.Post{}, err
	}

	if len(user.Posts) < 1 {
		return []models.Post{}, errors.New("no posts")
	}

	return user.Posts, nil
}

func GetUserFavoritesPaginated(userID uuid.UUID, resultsPerPage, pageNum uint) ([]models.Post, error) {

	user, err := repositories.GetUserByIdWithFavoritesPaginated(userID, resultsPerPage, pageNum)
	if err != nil {
		return []models.Post{}, err
	}

	if len(user.FavoritedPosts) < 1 {
		return []models.Post{}, errors.New("no favorite posts")
	}

	var postIDs []uuid.UUID
	for _, post := range user.FavoritedPosts {
		postIDs = append(postIDs, post.PostID)
	}

	posts, err := repositories.GetPostsByID(postIDs)
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}
