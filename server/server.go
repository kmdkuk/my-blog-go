package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/my-blog-go/controller"
	"github.com/kmdkuk/my-blog-go/infrastructure/mysql/persistence"
)

func SetUpRouting() *gin.Engine {
	r := gin.Default()
	base := r.Group("/api/v1")
	a := base.Group("/articles")
	articleRepository := persistence.NewArticleRepository()
	articleController := controller.NewArticleController(articleRepository)
	a.GET("", articleController.Index)
	a.GET("/:id", articleController.Show)
	a.POST("", articleController.Create)
	a.PUT("/:id", articleController.Update)
	a.DELETE("/:id", articleController.Delete)
	return r
}
