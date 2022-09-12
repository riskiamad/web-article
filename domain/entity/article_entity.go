package entity

import (
	"context"
	"time"

	"github.com/riskiamad/web-article/domain/model"
	"github.com/riskiamad/web-article/util"
)

// Article: struct which hold article entity
type Article struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleService: represent the article's service contract
type ArticleService interface {
	Store(ctx context.Context, articleCreateRequest *model.ArticleCreateRequest) (*model.ArticleResponse, error)
	Get(ctx context.Context, queryParams *util.QueryParams) ([]*model.ArticleResponse, error)
}

// ArticleRepository: represent the article's repository contract
type ArticleRepository interface {
	Store(ctx context.Context, article *Article) (*Article, error)
	GetAll(ctx context.Context, queryParams *util.QueryParams) ([]*model.ArticleResponse, error)
	GetWithFilter(ctx context.Context, queryParams *util.QueryParams) ([]*model.ArticleResponse, error)
}
