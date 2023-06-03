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
	s.router.GET("/ping", s.ping)
	s.router.GET("/get", s.getGeneralFollows)
}

func (s *server) ping(c *gin.Context) {
	response := s.service.Ping()
	c.JSON(response.StatusCode, response)
}

func (s *server) getGeneralFollows(c *gin.Context) {
	response := s.service.GetGeneralFollows(c.QueryArray("login"))
	c.JSON(response.StatusCode, response)
}
