package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/techvifyharrydo/Room-Reservation/ent"
	"github.com/techvifyharrydo/Room-Reservation/ent/office"
	"github.com/techvifyharrydo/Room-Reservation/internal/util"
)

type OfficeRepository interface {
	GetOffices(ctx context.Context) ([]*ent.Office, error)
	GetOffice(ctx context.Context, id uuid.UUID) (*ent.Office, error)
	CreateOffice(ctx context.Context, input ent.CreateOfficeInput) (*ent.Office, error)
	UpdateOffice(ctx context.Context, input ent.UpdateOfficeInput) (*ent.Office, error)
	DeleteOffice(ctx context.Context, id uuid.UUID) (error)
}

type officeRepoImpl struct {
	client *ent.Client
}

func NewOfficeRepository(client *ent.Client) OfficeRepository {
	return &officeRepoImpl{
		client: client,
	}
}

func (rps *officeRepoImpl) GetOffices(ctx context.Context) ([]*ent.Office, error) {
	return rps.client.Office.Query().All(ctx)
}

func (rps *officeRepoImpl) GetOffice(ctx context.Context, id uuid.UUID) (*ent.Office, error) {
	return rps.client.Office.Get(ctx, id)
}

func (rps *officeRepoImpl) CreateOffice(ctx context.Context, input ent.CreateOfficeInput) (*ent.Office, error) {
	exists, err := rps.client.Office.
		Query().
		Where(office.Name(input.Name)).
		Exist(ctx)
	if err != nil {
		return nil, util.WrapGQLError(ctx, "Office with this name already exists", 500, util.ErrorFlag(err.Error()))
	}

	if exists {
		return nil, util.WrapGQLError(ctx, "Office with this name already exists", 200, util.ErrorFlag("Invalid input provided"))
	}
	result, err := rps.client.Office.
						Create().
						SetName(input.Name).
						SetDescription(*input.Description).
						Save(ctx)
	if err != nil {
		return nil, err
	}					
	return result, nil
}

func (rps *officeRepoImpl) UpdateOffice(ctx context.Context, input ent.UpdateOfficeInput) (*ent.Office, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, err
	}
	exists, err := rps.client.Office.
		Query().
		Where(
			office.Name(*input.Name),
			office.IDNEQ(id),
		).
		Exist(ctx)
	if err != nil {
		return nil, util.WrapGQLError(ctx, "Error query office database", 500, util.ErrorFlag(err.Error()))
	}

	if exists {
		return nil, util.WrapGQLError(ctx, "Office with this name already exists", 200, "Invalid input provided")
	}

	office, err := rps.client.Office.
		UpdateOneID(id).
		SetName(*input.Name).
		SetDescription(*input.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return office, nil
}

func (rps *officeRepoImpl) DeleteOffice(ctx context.Context, id uuid.UUID) (error) {
	return rps.client.Office.DeleteOneID(id).Exec(ctx)
}
