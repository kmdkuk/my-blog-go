package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Show(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type articleController struct {
}

func NewArticleController() ArticleController {
	return &articleController{}
}

func (a articleController) Index(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a articleController) Create(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a articleController) Show(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a articleController) Update(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
func (a articleController) Delete(c *gin.Context) {
	c.String(http.StatusInternalServerError, "not implemented")
}
