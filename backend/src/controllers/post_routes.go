package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/ckpns/media-sharing-platform/backend/src/services"
	"github.com/ckpns/media-sharing-platform/backend/src/validators"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO: add proper JSON error responses

func PostRoutesRegister(router *gin.Engine) {

	postGroup := router.Group("/post")
	postGroup.POST("/", AuthRequired(), createPost)
	postGroup.GET("/:id", getPost)
	postGroup.DELETE("/:id", AuthRequired(), deletePost)
	postGroup.POST("/:id/favorite", AuthRequired(), favoritePost)
	postGroup.DELETE("/:id/favorite", AuthRequired(), unfavoritePost)
}

func createPost(context *gin.Context) {

	validator := validators.PostCreateValidator{}
	err := validator.Bind(context)
	if err != nil {
		context.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	userID, err := uuid.Parse(context.GetString("user_id"))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	file, err := context.FormFile("file")
	if err != nil {
		context.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	uid := uuid.NewString()
	ext := filepath.Ext(file.Filename) // TODO: file.Filename can't be trusted. Find another way to get extension.
	dst := fmt.Sprintf("./public/%s%s", uid, ext)

	err = context.SaveUploadedFile(file, dst)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	validator.Post.UserID = userID
	validator.Post.MediaPath = fmt.Sprintf("%s%s", uid, ext)

	err = services.CreatePost(&validator.Post)
	if err != nil {

		// TODO: delete file

		context.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"post_id":    validator.Post.ID.String(),
		"media_path": validator.Post.MediaPath,
	})
}

func getPost(context *gin.Context) {

	postID, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	post, err := services.GetPost(postID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusFound, gin.H{
		"post_id":     post.ID,
		"media_path":  post.MediaPath,
		"description": post.Description,
	})
}

func deletePost(context *gin.Context) {

	userID, err := uuid.Parse(context.GetString("user_id"))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	postID, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = services.DeletePost(postID, userID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func favoritePost(context *gin.Context) {

	userID, err := uuid.Parse(context.GetString("user_id"))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	postID, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = services.FavoritePost(postID, userID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func unfavoritePost(context *gin.Context) {

	userID, err := uuid.Parse(context.GetString("user_id"))
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	postID, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = services.UnfavoritePost(postID, userID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
