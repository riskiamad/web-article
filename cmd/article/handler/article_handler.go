package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/riskiamad/web-article/domain/entity"
	"github.com/riskiamad/web-article/domain/model"
	"github.com/riskiamad/web-article/util"
)

// ArticleHandler: struct hold httphandler for article
type ArticleHandler struct {
	articleService entity.ArticleService
}

// NewArticleHandler: initialize the article's resource endpoint
func NewArticleHandler(e *echo.Group, service entity.ArticleService) {
	handler := &ArticleHandler{
		articleService: service,
	}

	e.POST("/articles", handler.InsertArticle)
	e.GET("/articles", handler.GetArticles)
}

// InsertArticle: insert the article by given request body.
func (ah *ArticleHandler) InsertArticle(ctx echo.Context) (err error) {
	articleCreateRequest := new(model.ArticleCreateRequest)
	err = ctx.Bind(&articleCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := util.IsRequestValid(articleCreateRequest); !ok {
		return ctx.JSON(http.StatusBadRequest, util.CustomValidationError(err))
	}

	requestCtx := ctx.Request().Context()
	result, err := ah.articleService.Store(requestCtx, articleCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, result)
}

// GetArticles: get the articles by filter by given query param.
func (ah *ArticleHandler) GetArticles(ctx echo.Context) (err error) {
	params := ctx.QueryParams()

	queryParams := util.GetParamsValue(params)

	requestCtx := ctx.Request().Context()

	results, err := ah.articleService.Get(requestCtx, queryParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, results)
}
