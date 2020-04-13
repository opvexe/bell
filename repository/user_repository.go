package repository

import "github.com/sirupsen/logrus"

type UserRepository struct {
	logger *logrus.Logger
}

func NewUserRepository(logger *logrus.Logger) *UserRepository {
	return &UserRepository{logger:logger}
}


func (u *UserRepository) FindByUid(uid int)  {

}