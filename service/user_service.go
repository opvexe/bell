package service

import (
	"bell/repository"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepo *repository.UserRepository
	logger   *logrus.Logger
}

func NewUserService(userRepo *repository.UserRepository, logger *logrus.Logger) *UserService {
	return &UserService{userRepo: userRepo, logger: logger}
}

func (s *UserService) FindByUid(uid int) {
	s.userRepo.FindByUid(uid)
}
