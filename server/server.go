package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/my-blog-go/controller"
)

func SetUpRouting() *gin.Engine {
	r := gin.Default()
	base := r.Group("/api/v1")
	a := base.Group("/articles")
	articleController := controller.NewArticleController()
	a.GET("", articleController.Index)
	a.GET("/:id", articleController.Show)
	a.POST("", articleController.Create)
	a.PUT("/:id", articleController.Update)
	a.DELETE("/:id", articleController.Delete)
	return r
}
