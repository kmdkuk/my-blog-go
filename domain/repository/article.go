package repository

import "github.com/kmdkuk/my-blog-go/model"

type ArticleRepository interface {
	Insert(article model.Article) error
	FindAll() ([]model.Article, error)
	FindByID(ID string) (model.Article, error)
}
