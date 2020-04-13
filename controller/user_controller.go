package controller

import (
	"bell/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc *service.UserService
}

func NewUserController(userSvc *service.UserService) *UserController {
	return &UserController{userSvc:userSvc}
}

func (u *UserController) Login(c *gin.Context)  {

}
