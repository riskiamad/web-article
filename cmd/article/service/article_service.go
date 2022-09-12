package service

import (
	"context"
	"time"

	"github.com/riskiamad/web-article/domain/entity"
	"github.com/riskiamad/web-article/domain/model"
	"github.com/riskiamad/web-article/util"
)

type articleService struct {
	articleRepository entity.ArticleRepository
	timeoutCtx        time.Duration
}

// NewArticleService: create new articleService object as representation of ArticleService interface
func NewArticleService(repository entity.ArticleRepository, timeoutCtx time.Duration) entity.ArticleService {
	return &articleService{
		articleRepository: repository,
		timeoutCtx:        timeoutCtx,
	}
}

// Store: method to validate article's business logic before store into database
func (as *articleService) Store(ctx context.Context, articleCreateRequest *model.ArticleCreateRequest) (*model.ArticleResponse, error) {
	ctxDeadline, cancel := context.WithTimeout(ctx, as.timeoutCtx)
	defer cancel()

	article := &entity.Article{
		Author:    articleCreateRequest.Author,
		Title:     articleCreateRequest.Title,
		Body:      articleCreateRequest.Body,
		CreatedAt: time.Now(),
	}

	result, err := as.articleRepository.Store(ctxDeadline, article)
	if err != nil {
		return nil, err
	}

	articleResponse := &model.ArticleResponse{
		ID:        result.ID,
		Author:    result.Author,
		Title:     result.Title,
		Body:      result.Body,
		CreatedAt: result.CreatedAt,
	}

	return articleResponse, nil
}

// Get: get articles with filter by given query param.
func (as *articleService) Get(ctx context.Context, queryParams *util.QueryParams) ([]*model.ArticleResponse, error) {
	var (
		result []*model.ArticleResponse
		err    error
	)

	ctxDeadline, cancel := context.WithTimeout(ctx, as.timeoutCtx)
	defer cancel()

	if queryParams.Author == "" && queryParams.Query == "" {
		result, err = as.articleRepository.GetAll(ctxDeadline, queryParams)
		if err != nil {
			return nil, err
		}
	} else {
		result, err = as.articleRepository.GetWithFilter(ctxDeadline, queryParams)
		if err != nil {
			return nil, err
		}
	}

	return result, nil

}
