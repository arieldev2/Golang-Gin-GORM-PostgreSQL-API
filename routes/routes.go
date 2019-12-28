package routes

import (
	"gin_test/controller"

	"github.com/gin-gonic/gin"
)

func Routes() {

	r := gin.Default()
	r.GET("/", controller.AllPost)
	r.GET("/:id", controller.GetPost)
	r.POST("/", controller.NewPost)
	r.PUT("/:id", controller.UpdatePost)
	r.DELETE("/:id", controller.DeletePost)
	r.Run()
}
