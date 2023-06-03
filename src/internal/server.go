package internal

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type server struct {
	service *Service
	router  *gin.Engine
}

func NewServer(service Service) *server {
	return &server{router: gin.Default(), service: &service}
}

func (s *server) Start(port int) {
	s.initRouters()
	s.router.Run(":" + strconv.Itoa(port))
}

func (s *server) initRouters() {
	s.router.GET("/ping", s.service.Ping)
	s.router.GET("/get", s.service.GetGeneralFollows)
}
