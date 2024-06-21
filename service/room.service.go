package service

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/techvifyharrydo/Room-Reservation/ent"
	"github.com/techvifyharrydo/Room-Reservation/repository"
	"go.uber.org/zap"
)

type RoomService interface {
	GetRooms(ctx context.Context, filter *ent.RoomFilter) ([]*ent.Room, error)
	GetRoom(ctx context.Context, id uuid.UUID) (*ent.Room, error)
	CreateRoom(ctx context.Context, input ent.CreateRoomInput) (*ent.RoomResponse, error)
	UpdateRoom(ctx context.Context, input ent.UpdateRoomInput) (*ent.RoomResponse, error)
	DeleteRoom(ctx context.Context, id uuid.UUID) (string, error)
	GetAvailableRooms(ctx context.Context, input ent.GetAvailableRoomInput) ([]*ent.Room, error)
}

type roomSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewRoomService(repoRegistry repository.Repository, logger *zap.Logger) RoomService {
	return &roomSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *roomSvcImpl) GetRooms(ctx context.Context, filter *ent.RoomFilter) ([]*ent.Room, error) {
	return svc.repoRegistry.Room().GetRooms(ctx, *filter)
}

func (svc *roomSvcImpl) GetRoom(ctx context.Context, id uuid.UUID) (*ent.Room, error) {
	return svc.repoRegistry.Room().GetRoom(ctx, id)
}

func (svc *roomSvcImpl) CreateRoom(ctx context.Context, input ent.CreateRoomInput) (*ent.RoomResponse, error) {
	//validate
	if err := validateCreateRoomInput(input); err != nil {
		return nil, err
	}

	// create room in the database
	room, err := svc.repoRegistry.Room().CreateRoom(ctx, input)

	if err != nil {
		return nil, err
	}

	return &ent.RoomResponse{
		Message: "room is created ok",
		Data:    room,
	}, nil
}

func (svc *roomSvcImpl) UpdateRoom(ctx context.Context, input ent.UpdateRoomInput) (*ent.RoomResponse, error) {
	// validate
	if err := validateUpdateRoomInput(input); err != nil {
		return nil, err
	}

	// update room in the database
	room, err := svc.repoRegistry.Room().UpdateRoom(ctx, input)

	if err != nil {
		return nil, err
	}

	return &ent.RoomResponse{
		Message: "room is updated ok",
		Data:    room,
	}, nil
}

func (svc *roomSvcImpl) DeleteRoom(ctx context.Context, id uuid.UUID) (string, error) {
	// Validate id
	if id == uuid.Nil {
		return "", errors.New("invalid room ID")
	}

	err := svc.repoRegistry.Room().DeleteRoom(ctx, id)
	if err != nil {
		return "", err
	}

	return "Room deleted successfully", nil
}

func (svc *roomSvcImpl) GetAvailableRooms(ctx context.Context, input ent.GetAvailableRoomInput) ([]*ent.Room, error) {
	// Validate input
	if err := validateGetAvailableRoomInput(input); err != nil {
		return nil, err
	}

	return svc.repoRegistry.Room().GetAvailableRooms(ctx, input)
}

// Validation functions
func validateCreateRoomInput(input ent.CreateRoomInput) error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("room name is required")
	}
	if strings.TrimSpace(input.Color) == "" {
		return errors.New("room color is required")
	}
	if input.OfficeID == "" {
		return errors.New("office ID is required")
	}
	return nil
}

func validateUpdateRoomInput(input ent.UpdateRoomInput) error {
	if input.ID == uuid.Nil.String() {
		return errors.New("room ID is required")
	}
	if input.Name != nil && strings.TrimSpace(*input.Name) == "" {
		return errors.New("invalid room name")
	}
	if input.Color != nil && strings.TrimSpace(*input.Color) == "" {
		return errors.New("invalid room color")
	}
	if input.OfficeID != nil && *input.OfficeID == uuid.Nil.String() {
		return errors.New("invalid office ID")
	}
	return nil
}

func validateGetAvailableRoomInput(input ent.GetAvailableRoomInput) error {
	if input.StartDate == "" {
		return errors.New("start date is required")
	}
	if input.StartTime == "" {
		return errors.New("start time is required")
	}
	if input.EndTime == "" {
		return errors.New("end time is required")
	}
	if input.OfficeID == "" {
		return errors.New("office ID is required")
	}
	return nil
}
