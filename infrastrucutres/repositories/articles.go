package repositories

import (
	"go-serviceboilerplate/commons/models"
	"go-serviceboilerplate/commons/utils"
	"go-serviceboilerplate/infrastrucutres/databases/postgres"
)

type ArticlesRepository struct {
	postgresDB *postgres.PostgresInstance
}

func NewArticlesRepository(postgresDB *postgres.PostgresInstance) *ArticlesRepository {
	return &ArticlesRepository{postgresDB}
}

func (s *ArticlesRepository) CreateArticle(payload models.CreateArticleRequest) (string, error) {
	article := models.ArticleDTO{
		Title:   payload.Title,
		Content: payload.Content,
		Author: payload.Author,
		UpdatedBy: payload.Author,
	}

	result := s.postgresDB.DB.Create(&article)
	if result.Error != nil {
		return "", utils.NewInvariantError(result.Error)
	}

	return article.ID, nil
}

func (s *ArticlesRepository) GetArticles() ([]models.ArticleDTO, models.Metadata, error) {
	articles := []models.ArticleDTO{}
	metadata := models.Metadata{
	}
	
	result := s.postgresDB.DB.Find(&articles)
	if result.Error != nil {
		return articles, metadata, utils.NewInvariantError(result.Error)
	}

	metadata = models.Metadata{
		TotalCount:  len(articles),
		TotalPage:   1,
		CurrentPage: 1,
		PerPage:     len(articles),
	}

	return articles, metadata, nil
}