package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/my-blog-go/domain/repository"
)

type ArticleController interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Show(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type articleController struct {
	articleRepository repository.ArticleRepository
}

func NewArticleController(articleRepository repository.ArticleRepository) ArticleController {
	return &articleController{articleRepository: articleRepository}
}

func (a articleController) Index(c *gin.Context) {
	articles, _ := a.articleRepository.FindAll()
	c.JSON(http.StatusOK, articles)
}
func (a articleController) Create(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a articleController) Show(c *gin.Context) {
	id := c.Param("id")
	article, err := a.articleRepository.FindByID(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, article)
}
func (a articleController) Update(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a articleController) Delete(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
