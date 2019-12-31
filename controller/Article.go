package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var ArticleController Article

func init() {
	ArticleController = NewArticle()
}

type Article interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Show(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type article struct{}

func NewArticle() Article {
	return &article{}
}

func (a article) Index(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a article) Create(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a article) Show(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a article) Update(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a article) Delete(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
