package routes

import (
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/controllers/book_controller"
	"gin-gonic-gorm/controllers/file_controller"
	"gin-gonic-gorm/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
	//route user
	route.GET("/user", user_controller.GetAllUser)
	route.GET("/user/paginate", user_controller.UserPaginate)
	route.GET("/user/:id", user_controller.GetUserByID)
	route.POST("/user", user_controller.Store)
	route.PATCH("/user/:id", user_controller.Update)
	route.DELETE("/user/:id", user_controller.Delete)

	//route book
	route.GET("/book", book_controller.GetAllBook)

	//route book
	route.POST("/file", file_controller.HandleUploadFile)
}
