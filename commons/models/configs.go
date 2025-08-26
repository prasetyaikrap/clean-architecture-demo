package models

import "github.com/labstack/echo/v4/middleware"

type ApplicationConfig struct {
	Port 				string
	ClientID			string
	SecretToken			string
	JWTSecret			string
	CORSConfig			middleware.CORSConfig
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	TimeZone string
	SSLMode  string 
}

type ENVConfig struct {
	Application 	ApplicationConfig
	DB				DBConfig
}

type Metadata struct {
	TotalCount 	int `json:"total_count"`
	TotalPage  	int `json:"total_page"`
	CurrentPage int `json:"current_page"`
	PerPage    	int `json:"per_page"`
}