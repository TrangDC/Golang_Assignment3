package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/techvifyharrydo/Room-Reservation/ent"
)

type BookingRepository interface {
	// query
	GetBookings(ctx context.Context, filter ent.BookingFilter) ([]*ent.BookingData, error)
	GetBooking(ctx context.Context, id uuid.UUID) (*ent.BookingData, error)

	// mutation
	CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.Booking, error)
	UpdateBooking(ctx context.Context, input ent.UpdateBookingInput) (*ent.BookingResponse, error)
	CancelBooking(ctx context.Context, id uuid.UUID) (string, error)
}

type bookingRepoImpl struct {
	client *ent.Client
}

func NewBookingRepository(client *ent.Client) BookingRepository {
	return &bookingRepoImpl{
		client: client,
	}
}

func (rps *bookingRepoImpl) GetBookings(ctx context.Context, filter ent.BookingFilter) ([]*ent.BookingData, error) {
	return []*ent.BookingData{}, nil
}

func (rps *bookingRepoImpl) GetBooking(ctx context.Context, id uuid.UUID) (*ent.BookingData, error) {
	return &ent.BookingData{}, nil
}

func (rps *bookingRepoImpl) CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.Booking, error) {
	return nil, nil
}

func (rps *bookingRepoImpl) UpdateBooking(ctx context.Context, input ent.UpdateBookingInput) (*ent.BookingResponse, error) {
	return nil, nil
}

func (rps *bookingRepoImpl) CancelBooking(ctx context.Context, id uuid.UUID) (string, error) {
	return "", nil
}
