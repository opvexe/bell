package web

import (
	"bell/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *router.Router
}

func NewServer(engine *gin.Engine,apiRouter *router.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}

func (s *Server) Start()  {
	s.apiRouter.With(s.engine)
	s.engine.Run(":8080")
}
