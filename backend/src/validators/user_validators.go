package validators

import (
	"errors"
	"regexp"

	"github.com/ckpns/media-sharing-platform/backend/src/data/models"
	"github.com/gin-gonic/gin"
)

type UserRegisterValidator struct {
	RequestToValidate struct {
		Username string `json:"username" binding:"required,min=4,max=255"`
		Password string `json:"password" binding:"required,min=4,max=255"`
	}
	User models.User
}

var containsAtLeastOneUpperCaseLetter = regexp.MustCompile(`[A-Z]`)
var containsAtLeastOneLowerCaseLetter = regexp.MustCompile(`[a-z]`)
var containsAtLeastOneNumber = regexp.MustCompile(`[0-9]`)
var containsAtLeastOneSpecialCharacter = regexp.MustCompile(`[%#!$^&*@]`)

func (validator *UserRegisterValidator) Bind(context *gin.Context) error {

	if err := context.ShouldBindJSON(&validator.RequestToValidate); err != nil {
		return err
	}

	if !containsAtLeastOneUpperCaseLetter.MatchString(validator.RequestToValidate.Password) {
		return errors.New("password must contain an upper-case letter")
	}

	if !containsAtLeastOneLowerCaseLetter.MatchString(validator.RequestToValidate.Password) {
		return errors.New("password must contain a lower-case letter")
	}

	if !containsAtLeastOneNumber.MatchString(validator.RequestToValidate.Password) {
		return errors.New("password must contain at least one number")
	}

	if !containsAtLeastOneSpecialCharacter.MatchString(validator.RequestToValidate.Password) {
		return errors.New("password must contain at least one special character")
	}

	validator.User.Username = validator.RequestToValidate.Username
	validator.User.Password = validator.RequestToValidate.Password

	return nil
}

type UserLoginValidator struct {
	RequestToValidate struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user"`
	User models.User
}

func (validator *UserLoginValidator) Bind(context *gin.Context) error {

	if err := context.ShouldBindJSON(&validator.RequestToValidate); err != nil {
		return err
	}

	validator.User.Username = validator.RequestToValidate.Username
	validator.User.Password = validator.RequestToValidate.Password

	return nil
}
