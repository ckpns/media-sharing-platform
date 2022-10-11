package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ckpns/media-sharing-platform/backend/src/services"
	"github.com/ckpns/media-sharing-platform/backend/src/validators"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserRoutesRegister(router *gin.Engine) {

	userGroup := router.Group("/user")
	userGroup.POST("/register", registerUser)
	userGroup.POST("/login", loginUser)
	userGroup.DELETE("/", AuthRequired(), deleteUser)
	userGroup.GET("/posts/:page_num", AuthRequired(), getPostsPaginated)
	userGroup.GET("/favorites/:page_num", AuthRequired(), getFavoritesPaginated)
}

func registerUser(context *gin.Context) {

	validator := validators.UserRegisterValidator{}

	err := validator.Bind(context)
	if err != nil {
		context.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	err = services.CreateUser(&validator.User)
	if err != nil {

		// TODO: in case of sql duplicate key error, return user already exists

		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}

func loginUser(context *gin.Context) {

	validator := validators.UserLoginValidator{}

	err := validator.Bind(context)
	if err != nil {
		context.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	token, err := services.LoginUser(validator.User.Username, validator.User.Password)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func deleteUser(context *gin.Context) {

	userID, err := uuid.Parse(context.GetString("user_id"))
	fmt.Println(userID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = services.DeleteUser(userID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func getPostsPaginated(context *gin.Context) {

	userID, err := uuid.Parse(context.GetString("user_id"))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resultsPerPage, err := strconv.ParseUint(context.Query("results_per_page"), 10, 64)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	pageNum, err := strconv.ParseUint(context.Param("page_num"), 10, 64)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	posts, err := services.GetUserPostsPaginated(userID, uint(resultsPerPage), uint(pageNum))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Note: this is also hard-coded in getFavoritesPaginated
	// TODO: avoid duplicate implementation. Rewrite in a helpers package. (watch out for tight coupling)
	type postResponse struct {
		ID          string
		Description string
		MediaPath   string
	}
	var postsResponse []postResponse

	for _, post := range posts {
		postsResponse = append(postsResponse, postResponse{
			ID:          post.ID.String(),
			Description: post.Description,
			MediaPath:   post.MediaPath,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"posts": postsResponse,
	})
}

func getFavoritesPaginated(context *gin.Context) {

	userID, err := uuid.Parse(context.GetString("user_id"))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resultsPerPage, err := strconv.ParseUint(context.Query("results_per_page"), 10, 64)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	pageNum, err := strconv.ParseUint(context.Param("page_num"), 10, 64)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	posts, err := services.GetUserFavoritesPaginated(userID, uint(resultsPerPage), uint(pageNum))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	type postResponse struct {
		ID          string
		Description string
		MediaPath   string
	}
	var postsResponse []postResponse

	for _, post := range posts {
		postsResponse = append(postsResponse, postResponse{
			ID:          post.ID.String(),
			Description: post.Description,
			MediaPath:   post.MediaPath,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"posts": postsResponse,
	})
}
