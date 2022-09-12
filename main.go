package main

import (
	"time"

	"github.com/labstack/echo/v4"
	_articleHandler "github.com/riskiamad/web-article/cmd/article/handler"
	_articleRepository "github.com/riskiamad/web-article/cmd/article/repository"
	_articleService "github.com/riskiamad/web-article/cmd/article/service"
	"github.com/riskiamad/web-article/config"
	"github.com/riskiamad/web-article/database"
)

var (
	db         = database.DB
	env        = config.Config
	timeoutCtx = time.Duration(2 * time.Second)
	esClient   = database.EsClient
)

func main() {

	e := echo.New()

	// group router with prefix
	r := e.Group("/api/v1")

	// article dependencies
	articleRepository := _articleRepository.NewArticleRepository(db, esClient)
	articleService := _articleService.NewArticleService(articleRepository, timeoutCtx)
	_articleHandler.NewArticleHandler(r, articleService)

	// start server
	e.Logger.Fatal(e.Start(env.ServerHost))
}
