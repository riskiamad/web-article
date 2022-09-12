package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/olivere/elastic/v7"
	"github.com/riskiamad/web-article/domain/entity"
	"github.com/riskiamad/web-article/domain/model"
	"github.com/riskiamad/web-article/util"
)

type articleRepository struct {
	DB       *sql.DB
	EsClient *elastic.Client
}

// NewArticleRepository: create an object of ArticleRepository interface
func NewArticleRepository(db *sql.DB, esClient *elastic.Client) entity.ArticleRepository {
	return &articleRepository{
		DB:       db,
		EsClient: esClient,
	}
}

// Store: method for storing article data to database
func (ar *articleRepository) Store(ctx context.Context, article *entity.Article) (*entity.Article, error) {

	tx, err := ar.DB.Begin()
	if err != nil {
		return nil, err
	}

	rawQuery := "INSERT INTO article(author, title, body, created_at) values (author=?, title=?, body=?, created_at=?)"
	result, err := tx.ExecContext(ctx, rawQuery, article.Author, article.Title, article.Body, article.CreatedAt)
	if err != nil {
		println("ini bukan")
		tx.Rollback()
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	article.ID = lastID

	dataJSON, err := json.Marshal(&article)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	js := string(dataJSON)

	_, err = ar.EsClient.Index().
		Index("article").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return article, nil
}

// GetWithFilter: get articles from read database with filter by keyword or author
func (ar *articleRepository) GetWithFilter(ctx context.Context, queryParams *util.QueryParams) ([]*model.ArticleResponse, error) {
	var articles []*model.ArticleResponse

	authorQuery := elastic.NewMatchQuery("author", queryParams.Author)
	keywordQuery := elastic.NewMultiMatchQuery(queryParams.Query, "title", "body")

	qry := elastic.NewBoolQuery().Should(authorQuery).Should(keywordQuery)

	if queryParams.Author != "" {
		qry = elastic.NewBoolQuery().Must(authorQuery).Should(keywordQuery)
	}

	if queryParams.Author != "" && queryParams.Query != "" {
		qry = elastic.NewBoolQuery().Must(authorQuery).Must(keywordQuery)
	}

	sort := elastic.NewFieldSort(queryParams.OrderBy).Order(queryParams.Ascending)

	searchResult, err := ar.EsClient.Search().
		Index("article").
		Query(qry).
		SortBy(sort).
		From(queryParams.From).
		Size(queryParams.Perpage).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var article *model.ArticleResponse
		if err = json.Unmarshal(hit.Source, &article); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// GetAll: get articles from read database.
func (ar *articleRepository) GetAll(ctx context.Context, queryParams *util.QueryParams) ([]*model.ArticleResponse, error) {
	var articles []*model.ArticleResponse

	sort := elastic.NewFieldSort(queryParams.OrderBy).Order(queryParams.Ascending)

	searchResult, err := ar.EsClient.Search().
		Index("article").
		SortBy(sort).
		From(queryParams.From).
		Size(queryParams.Perpage).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var article *model.ArticleResponse
		if err = json.Unmarshal(hit.Source, &article); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}
