package articles

import (
	"fmt"
	"go-serviceboilerplate/interfaces/http/middlewares"

	"github.com/labstack/echo/v4"
)

func (h *ArticlesHandler) RegisterRoute(e *echo.Echo, middlewares *middlewares.AppMiddlewaresHandler) {
	g := e.Group("/articles")
	g.POST("", h.PostCreateArticle)
	g.GET("", h.GetArticleList)

	fmt.Println("Articles route registered")
}