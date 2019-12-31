package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/my-blog-go/controller"
)

func SetUpRouting() *gin.Engine {
	r := gin.Default()
	base := r.Group("/api/v1")
	a := base.Group("/articles")
	a.GET("", controller.ArticleController.Index)
	a.GET("/:id", controller.ArticleController.Show)
	a.POST("", controller.ArticleController.Create)
	a.PUT("/:id", controller.ArticleController.Update)
	a.DELETE("/:id", controller.ArticleController.Delete)
	return r
}
