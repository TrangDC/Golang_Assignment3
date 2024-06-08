package service

import (
	"context"
	"intern_traning/ent"
	"intern_traning/ent/news"
	"intern_traning/ent/predicate"
	"intern_traning/internal/util"
	"intern_traning/repository"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type NewsService interface {
	//mutation
	CreateNews(ctx context.Context, input ent.NewNewsInput) (*ent.NewsResponse, error)
	UpdateNews(ctx context.Context, id string, input ent.UpdateNewsInput) (*ent.NewsResponse, error)
	DeleteNews(ctx context.Context, id string) error

	//query
	GetNews(ctx context.Context, id string) (*ent.NewsResponse, error)
	GetAllNews(ctx context.Context, pagination *ent.PaginationInput, filter *ent.NewsFilter,
		orderBy *ent.NewsOrder, freeWord *ent.NewsFreeWord) (*ent.NewsResponseGetAll, error)
}

type newsSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewNewsService(repoRegistry repository.Repository, logger *zap.Logger) NewsService {
	return &newsSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

// mutation
func (svc newsSvcImpl) CreateNews(ctx context.Context, input ent.NewNewsInput) (*ent.NewsResponse, error) {
	var results *ent.News
	var err error
	err = svc.repoRegistry.News().ValidName(ctx, uuid.UUID{}, input.Title)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		results, err = repoRegistry.News().CreateNews(ctx, input)
		return err
	})
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	results, err = svc.repoRegistry.News().GetNews(ctx, results.ID)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.NewsResponse{
		Data: results,
	}, err
}

func (svc newsSvcImpl) UpdateNews(ctx context.Context, id string, input ent.UpdateNewsInput) (*ent.NewsResponse, error) {
	var results *ent.News
	var err error
	err = svc.repoRegistry.News().ValidName(ctx, uuid.MustParse(id), input.Title)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	record, err := svc.repoRegistry.News().GetNews(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		results, err = repoRegistry.News().UpdateNews(ctx, record, input)
		return err
	})
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	results, _ = svc.repoRegistry.News().GetNews(ctx, results.ID)
	return &ent.NewsResponse{
		Data: results,
	}, err
}

func (svc newsSvcImpl) DeleteNews(ctx context.Context, id string) error {
	var err error
	record, err := svc.repoRegistry.News().GetNews(ctx, uuid.MustParse(id))
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagCanNotCreate)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.News().DeleteNews(ctx, record)
		return err
	})
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return err
}

// query
func (svc newsSvcImpl) GetNews(ctx context.Context, id string) (*ent.NewsResponse, error) {
	record, err := svc.repoRegistry.News().GetNews(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.NewsResponse{
		Data: record,
	}, nil
}

func (svc newsSvcImpl) GetAllNews(ctx context.Context, pagination *ent.PaginationInput, filter *ent.NewsFilter,
	orderBy *ent.NewsOrder, freeWord *ent.NewsFreeWord) (*ent.NewsResponseGetAll, error) {
	var results *ent.NewsResponseGetAll
	var err error
	var page int
	var perPage int
	query := svc.repoRegistry.News().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	total, err := svc.repoRegistry.News().BuildCount(ctx, query)
	if err != nil {
		return results, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if orderBy != nil {
		if orderBy.Direction == ent.OrderDirectionAsc {
			query = query.Order(ent.Asc(strings.ToLower(orderBy.Field.String())))
		} else {
			query = query.Order(ent.Desc(strings.ToLower(orderBy.Field.String())))
		}
	} else {
		query = query.Order(ent.Desc(news.FieldCreatedAt))
	}
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(*pagination.PerPage).Offset((*pagination.Page - 1) * *pagination.PerPage)
	}
	records, err := svc.repoRegistry.News().BuildList(ctx, query)
	if err != nil {
		return results, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges := lo.Map(records, func(record *ent.News, index int) *ent.NewsEdge {
		return &ent.NewsEdge{
			Node: record,
			Cursor: ent.Cursor{
				Value: record.ID.String(),
			},
		}
	})
	return &ent.NewsResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   total,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

// common function
func (svc *newsSvcImpl) freeWord(newsQuery *ent.NewsQuery, input *ent.NewsFreeWord) {
	var predicateWhere []predicate.News
	if input != nil {
		if input.Title != nil {
			predicateWhere = append(predicateWhere, news.TitleContainsFold(*input.Title))
		}
		if input.Description != nil {
			predicateWhere = append(predicateWhere, news.DescriptionContainsFold(*input.Description))
		}
	}
	if len(predicateWhere) > 0 {
		newsQuery.Where(news.Or(predicateWhere...))
	}
}

func (svc *newsSvcImpl) filter(newsQuery *ent.NewsQuery, input *ent.NewsFilter) {
	if input != nil {
		if input.AuthorID != nil {
			newsQuery.Where(news.AuthorIDEQ(uuid.MustParse(*input.AuthorID)))
		}
		if input.Status != nil {
			newsQuery.Where(news.StatusEQ(news.Status(*input.Status)))
		}

		if input.FromDate != nil && input.ToDate != nil {
			newsQuery.Where(news.CreatedAtGTE(*input.FromDate), news.CreatedAtLTE(*input.ToDate))
		}
	}
}
