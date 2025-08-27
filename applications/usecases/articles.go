package usecases

import (
	"go-serviceboilerplate/commons/models"
	"go-serviceboilerplate/infrastrucutres/repositories"
)

type ArticleUsecases struct {
	articlesRepository *repositories.ArticlesRepository
}

func NewArticleUsecases(articlesRepository *repositories.ArticlesRepository) *ArticleUsecases {
	return &ArticleUsecases{articlesRepository}
}

func (r *ArticleUsecases) CreateArticle(payload models.CreateArticleRequest) (models.CreateArticleResponse, error) {	
	id, err := r.articlesRepository.CreateArticle(payload)

	response := models.CreateArticleResponse{ID: id}

	return response, err
}

func (r *ArticleUsecases) GetArticles() ([]models.GetArticlesResponse, models.Metadata, error) {	
	articlesResponse := []models.GetArticlesResponse{}
	articles, metadata, err := r.articlesRepository.GetArticles()

	if(err != nil) {
		return articlesResponse, metadata, err
	}

	for _, article := range articles {
		categories := []models.ArticleCategories{}
		article := models.GetArticlesResponse{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Author:    article.Author,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
			DeletedAt: article.DeletedAt,
			Categories: categories,
		}
		articlesResponse = append(articlesResponse, article)
	}

	return articlesResponse, metadata, err
}