package validators

import (
	"mime/multipart"

	"github.com/ckpns/media-sharing-platform/backend/src/data/models"
	"github.com/gin-gonic/gin"
)

type PostCreateValidator struct {
	RequestToValidate struct {
		File        *multipart.FileHeader `form:"file" binding:"required"`
		Description string                `form:"description" binding:"required,min=4,max=255"`
	}
	Post models.Post
}

func (validator *PostCreateValidator) Bind(context *gin.Context) error {

	if err := context.ShouldBind(&validator.RequestToValidate); err != nil {
		return err
	}

	validator.Post.Description = validator.RequestToValidate.Description

	return nil
}
