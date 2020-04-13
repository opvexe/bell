package web

import (
	"bell/config"
	"bell/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *router.Router
	logger    *logrus.Logger
	conf      *config.Config
}

func NewServer(engine *gin.Engine, apiRouter *router.Router, logger *logrus.Logger, conf *config.Config) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
		logger:    logger,
	}
}

func (s *Server) Start() {
	opt := s.conf.Http
	s.logger.Info(fmt.Sprintf("服务启动，运行模式：%s，版本号：%s，进程号：%d，端口号：%s", opt.RunModel, opt.Version, os.Getpid(), opt.Port))
	host := fmt.Sprintf("%s:%s", opt.Host, opt.Port)
	s.apiRouter.With(s.engine)
	if err := s.engine.Run(host); err != nil {
		panic(err)
	}
}
