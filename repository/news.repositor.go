package repository

import (
	"context"
	"fmt"
	"intern_traning/ent"
	"intern_traning/ent/news"
	"intern_traning/internal/util"
	"strings"
	"time"

	"github.com/google/uuid"
)

type NewsRepository interface {
	//mutation
	CreateNews(ctx context.Context, input ent.NewNewsInput) (*ent.News, error)
	UpdateNews(ctx context.Context, model *ent.News, input ent.UpdateNewsInput) (*ent.News, error)
	DeleteNews(ctx context.Context, model *ent.News) (*ent.News, error)

	BuildUpdateOne(ctx context.Context, model *ent.News) *ent.NewsUpdateOne
	BuildSaveUpdateOne(ctx context.Context, update *ent.NewsUpdateOne) (*ent.News, error)

	//query
	BuildQuery() *ent.NewsQuery
	BuildCount(ctx context.Context, query *ent.NewsQuery) (int, error)
	BuildList(ctx context.Context, query *ent.NewsQuery) ([]*ent.News, error)
	BuildGet(ctx context.Context, query *ent.NewsQuery) (*ent.News, error)
	GetNews(ctx context.Context, id uuid.UUID) (*ent.News, error)

	// common
	ValidName(ctx context.Context, newsId uuid.UUID, title string) error
}

type newsRepoImpl struct {
	client *ent.Client
}

func NewNewsRepository(client *ent.Client) NewsRepository {
	return &newsRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *newsRepoImpl) BuildCreate() *ent.NewsCreate {
	return rps.client.News.Create().SetUpdatedAt(time.Now())
}

func (rps *newsRepoImpl) BuildUpdate() *ent.NewsUpdate {
	return rps.client.News.Update().SetUpdatedAt(time.Now())
}

func (rps *newsRepoImpl) BuildDelete() *ent.NewsUpdate {
	return rps.client.News.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *newsRepoImpl) BuildQuery() *ent.NewsQuery {
	return rps.client.News.Query().Where(news.DeletedAtIsNil()).WithAuthorEdge()
}

func (rps *newsRepoImpl) BuildGet(ctx context.Context, query *ent.NewsQuery) (*ent.News, error) {
	return query.First(ctx)
}

func (rps *newsRepoImpl) BuildList(ctx context.Context, query *ent.NewsQuery) ([]*ent.News, error) {
	return query.All(ctx)
}

func (rps *newsRepoImpl) BuildCount(ctx context.Context, query *ent.NewsQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *newsRepoImpl) BuildExist(ctx context.Context, query *ent.NewsQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *newsRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.News) *ent.NewsUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *newsRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.NewsUpdateOne) (*ent.News, error) {
	return update.Save(ctx)
}

// mutation
func (rps *newsRepoImpl) CreateNews(ctx context.Context, input ent.NewNewsInput) (*ent.News, error) {
	oid := ctx.Value("user_id").(uuid.UUID)
	create := rps.BuildCreate().
		SetTitle(strings.TrimSpace(input.Title)).
		SetSlug(util.SlugGeneration(input.Title)).
		SetDescription(strings.TrimSpace(input.Description)).
		SetContent(strings.TrimSpace(input.Content)).
		SetAuthorID(oid)
	switch news.Status(input.Status) {
	case news.StatusPublished:
		create.SetStatus(news.StatusPublished)
	case news.StatusHidden:
		create.SetStatus(news.StatusHidden)
	default:
		create.SetStatus(news.StatusDraft)
	}
	return create.Save(ctx)
}

func (rps *newsRepoImpl) UpdateNews(ctx context.Context, model *ent.News, input ent.UpdateNewsInput) (*ent.News, error) {
	update := rps.BuildUpdateOne(ctx, model).SetTitle(strings.TrimSpace(input.Title)).
		SetSlug(util.SlugGeneration(input.Title)).
		SetDescription(strings.TrimSpace(input.Description)).
		SetContent(strings.TrimSpace(input.Content))
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *newsRepoImpl) DeleteNews(ctx context.Context, model *ent.News) (*ent.News, error) {
	update := rps.BuildUpdateOne(ctx, model).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
	return update.Save(ctx)
}

// query
func (rps *newsRepoImpl) GetNews(ctx context.Context, id uuid.UUID) (*ent.News, error) {
	query := rps.BuildQuery().Where(news.IDEQ(id))
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *newsRepoImpl) ValidName(ctx context.Context, newsId uuid.UUID, title string) error {
	query := rps.BuildQuery().Where(news.SlugEQ(util.SlugGeneration(title)))
	if newsId != uuid.Nil {
		query = query.Where(news.IDNEQ(newsId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("module.news.validation.title_exist")
	}
	return nil
}
