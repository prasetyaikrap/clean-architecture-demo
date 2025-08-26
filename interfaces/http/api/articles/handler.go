package articles

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-serviceboilerplate/applications/usecases"
	"go-serviceboilerplate/commons/models"
	"go-serviceboilerplate/interfaces/utils"
)

type ArticlesHandler struct {
	articleUsecases *usecases.ArticleUsecases
}

func NewArticlesHandler(articleUsecases *usecases.ArticleUsecases ) *ArticlesHandler {
	return &ArticlesHandler{articleUsecases}
}

// CreateArticle godoc
// @Summary      Create an article
// @Description  Add a new article to the database
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        article  body      models.CreateArticleRequest  true  "Article"
// @Success      200      {object}  utils.SuccessResponseConfig
// @Failure      400      {object}  utils.ErrorResponseConfig
// @Router       /articles [post]
func (h *ArticlesHandler) PostCreateArticle(c echo.Context) error {
	request := models.CreateArticleRequest{}
	if err := c.Bind(&request); err != nil {
		return utils.ErrorResponse(c, err)
	}

	if err := c.Validate(&request); err != nil {
		return utils.ErrorResponse(c, err) 
	}

	data, err := h.articleUsecases.CreateArticle(request)
	
	if(err != nil) {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, utils.SuccessResponseConfig{
		Code:    http.StatusCreated,
		Message: "Article Created successfully",
		Data:    data,
	})
}

func (h *ArticlesHandler) GetArticleList(c echo.Context) error {
	data, metadata, err := h.articleUsecases.GetArticles()
	
	if(err != nil) {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponseWithMetadata(c, utils.SuccessResponseWithMetadataConfig{
		Code:    http.StatusOK,
		Message: "Article Retrieved successfully",
		Data:    data,
		Metadata: metadata,
	})
}