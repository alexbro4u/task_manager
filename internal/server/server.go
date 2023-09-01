package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	r := gin.Default()
	return &Server{
		router: r,
	}
}

func (s *Server) Start(port string) error {
	return s.router.Run(":" + port)
}
