package service

import (
	"github.com/techvifyharrydo/Room-Reservation/ent"
	"github.com/techvifyharrydo/Room-Reservation/internal/azuread"
	"github.com/techvifyharrydo/Room-Reservation/repository"
	"go.uber.org/zap"
)

type Service interface {
	Auth() AuthService
	User() UserService
	Office() OfficeService
	Room() RoomService
	Booking() BookingService
}

type serviceImpl struct {
	authService    AuthService
	UserService    UserService
	OfficeService  OfficeService
	RoomService    RoomService
	BookingService BookingService
}

func NewService(azureADOAuthClient azuread.AzureADOAuth, entClient *ent.Client, logger *zap.Logger) Service {
	repoRegistry := repository.NewRepository(entClient)

	return &serviceImpl{
		authService:    NewAuthService(azureADOAuthClient, logger),
		UserService:    NewUserService(repoRegistry, logger),
		OfficeService:  NewOfficeService(repoRegistry, logger),
		RoomService:    NewRoomService(repoRegistry, logger),
		BookingService: NewBookingService(repoRegistry, logger),
	}
}

func (i serviceImpl) Auth() AuthService {
	return i.authService
}

func (i serviceImpl) User() UserService {
	return i.UserService
}

func (i serviceImpl) Office() OfficeService {
	return i.OfficeService
}

func (i serviceImpl) Room() RoomService {
	return i.RoomService
}

func (i serviceImpl) Booking() BookingService {
	return i.BookingService
}
