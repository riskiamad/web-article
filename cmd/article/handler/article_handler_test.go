package handler_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/olivere/elastic/v7"
	"github.com/riskiamad/web-article/config"
	"github.com/riskiamad/web-article/domain/entity"
	"github.com/riskiamad/web-article/domain/model"
	"github.com/stretchr/testify/assert"

	_articleHandler "github.com/riskiamad/web-article/cmd/article/handler"
	_articleRepository "github.com/riskiamad/web-article/cmd/article/repository"
	_articleService "github.com/riskiamad/web-article/cmd/article/service"
)

var (
	env        = config.Config
	esClient   = newTestClient()
	dbConn     = connTestDB()
	timeoutCtx = time.Duration(2 * time.Second)
)

// connTestDB: MySQL connection for test.
func connTestDB() *sql.DB {
	db, err := sql.Open("mysql", env.DbUser+":"+env.DbPass+"@tcp("+env.DbHost+")/"+env.DbTestName)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// newTestClient: elasticsearch client for test.
func newTestClient() *elastic.Client {
	esClient, err := elastic.NewClient(
		elastic.SetURL(env.EsTestURL),
		elastic.SetSniff(false),
	)

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	return esClient
}

// setupTestRouter: mock of router for testing.
func setupTestRouter(dbConn *sql.DB, esClient *elastic.Client, timeoutCtx time.Duration) *echo.Echo {
	e := echo.New()

	// group router with prefix
	r := e.Group("/api/v1")

	// article dependencies
	articleRepository := _articleRepository.NewArticleRepository(dbConn, esClient)
	articleService := _articleService.NewArticleService(articleRepository, timeoutCtx)
	_articleHandler.NewArticleHandler(r, articleService)

	return e
}

func TestInsertArticleSuccess(t *testing.T) {
	router := setupTestRouter(dbConn, esClient, timeoutCtx)

	requestBody := strings.NewReader(`{
		"author": "bjorka",
		"title": "Cara mudah megatasi data bocor",
		"body": "Pakai No Drop!"
		}`)

	request := httptest.NewRequest(echo.POST, env.ServerHost+"/api/v1/articles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody model.ArticleResponse
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "bjorka", responseBody.Author)
	assert.Equal(t, "Cara mudah megatasi data bocor", responseBody.Title)
	assert.Equal(t, "Pakai No Drop!", responseBody.Body)
}

func TestInsertArticleFailed(t *testing.T) {
	router := setupTestRouter(dbConn, esClient, timeoutCtx)

	requestBody := strings.NewReader(`{
		"author": "bjorka",
		"title": "Cara mudah megatasi data bocor"
		}`)

	request := httptest.NewRequest(echo.POST, env.ServerHost+"/api/v1/articles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestGetArticlesSuccess(t *testing.T) {
	router := setupTestRouter(dbConn, esClient, timeoutCtx)

	repository := _articleRepository.NewArticleRepository(dbConn, esClient)
	article, _ := repository.Store(context.Background(), &entity.Article{
		Author:    "Fery",
		Title:     "Cerita hoaks",
		Body:      "Sebuah cerita yang di buat hanya untuk viral dan cari sensasi",
		CreatedAt: time.Now(),
	})

	request := httptest.NewRequest(echo.GET, env.ServerHost+"/api/v1/articles", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody []model.ArticleResponse
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, article.Author, responseBody[0].Author)
	assert.Equal(t, article.Title, responseBody[0].Title)
	assert.Equal(t, article.Body, responseBody[0].Body)
	assert.Equal(t, article.CreatedAt, responseBody[0].CreatedAt)
}
