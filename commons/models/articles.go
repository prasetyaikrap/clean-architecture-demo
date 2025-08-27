package models

import "time"

type CreateArticleRequest struct {
	Title      				string   	`json:"title" validate:"required,min=3"`
	Content    				string   	`json:"content" validate:"required,min=10"`
	Author      			string   	`json:"author" validate:"required,min=3"`
}

type CreateArticleResponse struct {
	ID       				string   	`json:"id"`
}

type GetArticlesResponse struct {
	ID        				string       		`json:"id"`
	Title					string				`json:"title"`
	Content					string				`json:"content"`
	Author					string				`json:"author"`
	CreatedAt 				time.Time			`json:"created_at"`
	UpdatedAt 				time.Time			`json:"updated_at"`
	DeletedAt 				time.Time 			`json:"deleted_at"`
	Categories 				[]ArticleCategories `json:"categories"`
}

type ArticleCategories struct {
	ID        				string          `json:"id"`
	Title					string			`json:"title"`
	Description				string			`json:"description"`
}