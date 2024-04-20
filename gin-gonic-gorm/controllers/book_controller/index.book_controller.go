package book_controller

import "github.com/gin-gonic/gin"

func GetAllBook(ctx *gin.Context) {
	isValidated := true

	if !isValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "something wrong with your request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"hello": "book",
	})
}
