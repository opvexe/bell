package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewUserRepository(logger *logrus.Logger, db *gorm.DB) *UserRepository {
	return &UserRepository{logger: logger, db: db}
}

func (u *UserRepository) FindByUid(uid int) {

}
