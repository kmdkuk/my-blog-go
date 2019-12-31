package persistence

import (
	"time"

	"github.com/kmdkuk/my-blog-go/domain/repository"
	"github.com/kmdkuk/my-blog-go/model"
)

type articleRepository struct{}

func NewArticleRepository() repository.ArticleRepository {
	return &articleRepository{}
}

func (ar *articleRepository) Insert(article model.Article) error {
	now := time.Now()
	article.CreatedAt = now
	article.UpdatedAt = now
	if err := DB.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) FindAll() ([]model.Article, error) {
	var articles []model.Article
	if err := DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (ar *articleRepository) FindByID(ID string) (model.Article, error) {
	var article model.Article
	if err := DB.Where("ID = ?", ID).First(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}
