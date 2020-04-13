package web

import (
	"bell/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *router.Router
	logger  *logrus.Logger
}

func NewServer(engine *gin.Engine,apiRouter *router.Router,logger *logrus.Logger) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
		logger:logger,
	}
}

func (s *Server) Start()  {
	s.apiRouter.With(s.engine)
	s.engine.Run(":8080")
}
