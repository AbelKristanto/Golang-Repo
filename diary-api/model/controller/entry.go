package controller

import (
	"diary-api/model"
	"diary-api/model/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddEntry(ctx *gin.Context) {
	var input model.Entry
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := helper.CurrentUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": savedEntry,
	})
}

func GetAllEntries(ctx *gin.Context) {
	user, err := helper.CurrentUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user.Entries,
	})
}
