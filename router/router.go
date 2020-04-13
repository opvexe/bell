package router

import (
	ctrls "bell/controller"
	"bell/middleware"
	"github.com/gin-gonic/gin"
	_ "bell/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Router struct {
	recoverMid *middleware.Recover

	userCtrl *ctrls.UserController
}

func NewRouter(recoverMid *middleware.Recover, userCtrl *ctrls.UserController) *Router {
	return &Router{recoverMid: recoverMid, userCtrl: userCtrl}
}

func (r *Router) With(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.Use(r.recoverMid.Handler())

	user := r.newUser(v1)
	{
		user.GET("/register", r.userCtrl.Create)
	}
	engine.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *Router) newUser(v1 *gin.RouterGroup) *gin.RouterGroup {
	public := v1.Group("/")
	// 添加中间件
	return public
}
