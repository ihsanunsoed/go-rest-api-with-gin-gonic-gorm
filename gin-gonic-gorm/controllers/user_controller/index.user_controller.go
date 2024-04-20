package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {

	users := new([]models.User)
	// database.DB.Find(&users)
	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error.",
		})

		return
	}
	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})

		return
	}
	ctx.JSON(200, gin.H{
		"message": "data transmitted.",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	emailExist := new(models.User)

	database.DB.Table("users").Where("email = ?", userReq.Email).First(&emailExist)

	if emailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
		})
		return
	}
	user := new(models.User)
	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errDb := database.DB.Table("users").Create(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Can't create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Data created successfully",
		"Data":    user,
	})
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)
	userEmailExist := new(models.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "data not found",
		})
		return
	}
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error

	if errUserEmailExist != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(400, gin.H{
			"message": "Email already exist",
		})
		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "cant update data",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data updated successfully.",
		"data":    user,
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)

	errUser := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errUser != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error.",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "data not found.",
		})
		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.User{}).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error.",
			"error":   errDb.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "data deleted successfully.",
	})
}

func UserPaginate(ctx *gin.Context) {

	pages := ctx.Query("pages")

	if pages == "" {
		pages = "1"
	}

	perPages := ctx.Query("perPages")

	if perPages == "" {
		perPages = "10"
	}

	perPageInt, _ := strconv.Atoi(perPages)
	pageInt, _ := strconv.Atoi(pages)

	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]models.User)

	// database.DB.Find(&users)
	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}
