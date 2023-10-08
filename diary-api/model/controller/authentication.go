package controller

import (
	"diary-api/model"
	"diary-api/model/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	var input model.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	//passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	//
	//if err != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//}
	//
	//user := model.User{
	//	Username: input.Username,
	//	Password: string(passwordHash),
	//}

	beforesavedUser, err := user.BeforeSave()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//savedUser, err := user.Save()
	savedUser, err := beforesavedUser.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	jwt, err := helper.GenerateJWT(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"jwt": jwt,
	})
}
