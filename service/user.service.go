package service

import (
	"github.com/techvifyharrydo/Room-Reservation/repository"
	"go.uber.org/zap"
)

type UserService interface {
	PreFunc() error
}

type userSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewUserService(repoRegistry repository.Repository, logger *zap.Logger) UserService {
	return &userSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *userSvcImpl) PreFunc() error {
	return nil
}
