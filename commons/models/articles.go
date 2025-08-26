package models

import (
	"time"
)

type ArticleDTO struct {
	ID        				uint           	`json:"id" gorm:"primaryKey"`
	Title					string			`json:"title" gorm:"type:varchar(255);not null"`
	Content					string			`json:"content" gorm:"type:text;not null"`
	Author					string			`json:"author" gorm:"type:varchar(100);not null"`
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`
	DeletedAt 				time.Time 		`json:"deleted_at" gorm:"index"`

	Categories 				[]CategoriesDTO `json:"categories" gorm:"many2many:articles_categories;"`
}

func (ArticleDTO) TableName() string {
	return "articles"
}

type CategoriesDTO struct {
	ID        				uint           	`json:"id" gorm:"primaryKey"`
	Title					string			`json:"title" gorm:"type:varchar(255);not null"`
	Description				string			`json:"description" gorm:"type:text;not null"`
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`
	DeletedAt 				time.Time 		`json:"deleted_at" gorm:"index"`

	Articles []ArticleDTO `json:"articles" gorm:"many2many:articles_categories;"`
}

func (CategoriesDTO) TableName() string {
	return "categories"
}