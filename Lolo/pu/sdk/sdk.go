package sdk

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New(router *gin.Engine) *Server {
	s := &Server{
		router: router,
	}
	s.Router()
	return s
}

func (s *Server) Close() {}
