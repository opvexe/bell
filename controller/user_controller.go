package controller

import (
	"bell/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	logger  *logrus.Logger
	userSvc *service.UserService
}

func NewUserController(logger *logrus.Logger, userSvc *service.UserService) *UserController {
	return &UserController{
		logger:  logger,
		userSvc: userSvc,
	}
}

func (u *UserController) Login(c *gin.Context) {

}
