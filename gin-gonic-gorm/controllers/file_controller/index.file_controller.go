package file_controller

import (
	"fmt"
	"gin-gonic-gorm/utils"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleUploadFile(ctx *gin.Context) {

	fileheader, _ := ctx.FormFile("file")

	if fileheader == nil {
		ctx.JSON(400, gin.H{
			"message": "file required",
		})
		return
	}

	extensionFile := filepath.Ext(fileheader.Filename)
	currentTime := time.Now().UTC().Format("20061206")
	filename := fmt.Sprintf("%s-%s%s", currentTime, utils.RandomString(16), extensionFile)
	errUpload := ctx.SaveUploadedFile(fileheader, fmt.Sprintf("./public/files/%s", filename))

	if errUpload != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error.",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "file uploaded successfuly.",
	})
}
