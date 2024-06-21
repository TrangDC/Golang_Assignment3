package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/techvifyharrydo/Room-Reservation/ent"
	"github.com/techvifyharrydo/Room-Reservation/internal/util"
	"github.com/techvifyharrydo/Room-Reservation/repository"
	"go.uber.org/zap"
)

type OfficeService interface {
	GetOffices(ctx context.Context) ([]*ent.Office, error)
	GetOffice(ctx context.Context, id uuid.UUID) (*ent.Office, error)
	CreateOffice(ctx context.Context, input ent.CreateOfficeInput) (*ent.OfficeResponse, error)
	UpdateOffice(ctx context.Context, input ent.UpdateOfficeInput) (*ent.OfficeResponse, error)
	DeleteOffice(ctx context.Context, id uuid.UUID) (string, error)
}

type officeSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewOfficeService(repoRegistry repository.Repository, logger *zap.Logger) OfficeService {
	return &officeSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *officeSvcImpl) GetOffices(ctx context.Context) ([]*ent.Office, error) {
	offices, err:= svc.repoRegistry.Office().GetOffices(ctx)
	if err!= nil {
        return nil, util.WrapGQLInternalError(ctx)
    }
	return offices, nil
}

func (svc *officeSvcImpl) GetOffice(ctx context.Context, id uuid.UUID) (*ent.Office, error) {
	return svc.repoRegistry.Office().GetOffice(ctx, id)
}

func (svc *officeSvcImpl) CreateOffice(ctx context.Context, input ent.CreateOfficeInput) (*ent.OfficeResponse, error) {
	office, err := svc.repoRegistry.Office().CreateOffice(ctx, input)
	if err != nil {
		return nil, util.WrapGQLError(ctx, "Failed to create office.", 500, util.ErrorFlag(err.Error()))
	}
	return &ent.OfficeResponse{
		Message: "Data has been successfully created.",
		Data:    office,
	}, err
}

func (svc *officeSvcImpl) UpdateOffice(ctx context.Context, input ent.UpdateOfficeInput) (*ent.OfficeResponse, error) {
	officeUpdate, err := svc.repoRegistry.Office().UpdateOffice(ctx, input)
	if err != nil {
		return nil, util.WrapGQLError(ctx, "Failed to update office.", 500, util.ErrorFlag(err.Error()))
	}
	return &ent.OfficeResponse{
		Message: "Data has been successfully updated.",
		Data:    officeUpdate,
	}, err
}

func (svc *officeSvcImpl) DeleteOffice(ctx context.Context, id uuid.UUID) (string, error) {
	errormessage := svc.repoRegistry.Office().DeleteOffice(ctx, id)
	if errormessage != nil {
		return "", util.WrapGQLError(ctx, "Failed to delete office.", 500, util.ErrorFlag(errormessage.Error()))
	}
	return "Data has been successfully deleted.", nil
}
