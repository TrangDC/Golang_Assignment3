package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/techvifyharrydo/Room-Reservation/ent"
	"github.com/techvifyharrydo/Room-Reservation/ent/booking"
	"github.com/techvifyharrydo/Room-Reservation/ent/room"
)

type RoomRepository interface {
	GetRooms(ctx context.Context, filter ent.RoomFilter) ([]*ent.Room, error)
	GetRoom(ctx context.Context, id uuid.UUID) (*ent.Room, error)
	CreateRoom(ctx context.Context, input ent.CreateRoomInput) (*ent.Room, error)
	UpdateRoom(ctx context.Context, input ent.UpdateRoomInput) (*ent.Room, error)
	DeleteRoom(ctx context.Context, id uuid.UUID) error
	GetAvailableRooms(ctx context.Context, input ent.GetAvailableRoomInput) ([]*ent.Room, error)
}

type roomRepoImpl struct {
	client *ent.Client
}

func NewRoomRepository(client *ent.Client) RoomRepository {
	return &roomRepoImpl{
		client: client,
	}
}

func (rps *roomRepoImpl) GetRooms(ctx context.Context, filter ent.RoomFilter) ([]*ent.Room, error) {
	filterOfficeID, err := uuid.Parse(filter.OfficeID)
	if err != nil {
		return nil, err
	}

	return rps.client.Room.Query().Where(room.OfficeID(filterOfficeID)).All(ctx)
}

func (rps *roomRepoImpl) GetRoom(ctx context.Context, id uuid.UUID) (*ent.Room, error) {
	return rps.client.Room.Get(ctx, id)
}

func (rps *roomRepoImpl) CreateRoom(ctx context.Context, input ent.CreateRoomInput) (*ent.Room, error) {
	officeID, err := uuid.Parse(input.OfficeID)
	if err != nil {
		return nil, err
	}
	return rps.client.Room.Create().
		SetName(input.Name).
		SetColor(input.Color).
		SetOfficeID(officeID).
		SetDescription(*input.Description).
		SetImageURL(*input.ImageURL).
		Save(ctx)
}

func (rps *roomRepoImpl) UpdateRoom(ctx context.Context, input ent.UpdateRoomInput) (*ent.Room, error) {
	roomID, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, err
	}

	upd := rps.client.Room.UpdateOneID(roomID)
	if input.Name != nil {
		upd.SetName(*input.Name)
	}
	if input.Color != nil {
		upd.SetColor(*input.Color)
	}
	if input.Description != nil {
		upd.SetDescription(*input.Description)
	}
	if input.ImageURL != nil {
		upd.SetImageURL(*input.ImageURL)
	}

	if input.OfficeID != nil {
		officeID, err := uuid.Parse(*input.OfficeID)
		if err != nil {
			return nil, err
		}
		upd.SetOfficeID(officeID)
	}

	return upd.Save(ctx)
}

func (rps *roomRepoImpl) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	return rps.client.Room.DeleteOneID(id).Exec(ctx)
}

func (rps *roomRepoImpl) GetAvailableRooms(ctx context.Context, input ent.GetAvailableRoomInput) ([]*ent.Room, error) {
	// Parse input times
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return nil, err
	}
	startTime, err := time.Parse("15:04", input.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("15:04", input.EndTime)
	if err != nil {
		return nil, err
	}

	startDateTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), startTime.Hour(), startTime.Minute(), 0, 0, time.UTC)
	endDateTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), endTime.Hour(), endTime.Minute(), 0, 0, time.UTC)

	// Retrieve all rooms in the specified office
	officeID, err := uuid.Parse(input.OfficeID)
	if err != nil {
		return nil, err
	}
	rooms, err := rps.client.Room.Query().Where(room.OfficeID(officeID)).All(ctx)
	if err != nil {
		return nil, err
	}

	// Find rooms with no overlapping bookings
	var availableRooms []*ent.Room
	for _, rm := range rooms {
		// Check for any overlapping bookings for this room
		conflictingBookings, err := rps.client.Booking.Query().
			Where(
				booking.RoomID(rm.ID),
				booking.Or(
					booking.And(
						booking.StartDateLTE(startDateTime),
						booking.EndDateGTE(endDateTime),
					),
					booking.And(
						booking.StartDateGTE(startDateTime),
						booking.StartDateLT(endDateTime),
					),
					booking.And(
						booking.EndDateGT(startDateTime),
						booking.EndDateLTE(endDateTime),
					),
				),
			).All(ctx)
		if err != nil {
			return nil, err
		}

		// If no conflicting bookings, add to available rooms
		if len(conflictingBookings) == 0 {
			availableRooms = append(availableRooms, rm)
		}
	}

	return availableRooms, nil
}
