package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/techvifyharrydo/Room-Reservation/ent"
	"github.com/techvifyharrydo/Room-Reservation/repository"
	"go.uber.org/zap"
)

type BookingService interface {
	// queries
	GetBookings(ctx context.Context, filter ent.BookingFilter) ([]*ent.BookingData, error)
	GetBooking(ctx context.Context, id uuid.UUID) (*ent.BookingData, error)

	// mutations
	CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.BookingResponse, error)
	UpdateBooking(ctx context.Context, input ent.UpdateBookingInput) (*ent.BookingResponse, error)
	CancelBooking(ctx context.Context, id uuid.UUID) (string, error)
}

type bookingSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewBookingService(repoRegistry repository.Repository, logger *zap.Logger) BookingService {
	return &bookingSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *bookingSvcImpl) GetBookings(ctx context.Context, filter ent.BookingFilter) ([]*ent.BookingData, error) {
	return []*ent.BookingData{}, nil
}

func (svc *bookingSvcImpl) GetBooking(ctx context.Context, id uuid.UUID) (*ent.BookingData, error) {
	return &ent.BookingData{}, nil
}

func (svc *bookingSvcImpl) CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.BookingResponse, error) {
	return &ent.BookingResponse{
		Data:    &ent.BookingData{},
		Message: "Data has been successfully created.",
	}, nil
}

func (svc *bookingSvcImpl) UpdateBooking(ctx context.Context, input ent.UpdateBookingInput) (*ent.BookingResponse, error) {
	return &ent.BookingResponse{
		Data:    &ent.BookingData{},
		Message: "Data has been successfully updated.",
	}, nil
}

func (svc *bookingSvcImpl) CancelBooking(ctx context.Context, id uuid.UUID) (string, error) {
	_, err := svc.repoRegistry.Booking().CancelBooking(ctx, id)
	if err != nil {
		return "", err
	}
	return "Data has been successfully deleted.", err
}
