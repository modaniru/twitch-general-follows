package internal

import "github.com/gin-gonic/gin"

type server struct {
	Config *configauration
	Router *gin.Engine
}

func NewServer(configPath string) *server {
	return &server{NewConfiguration(configPath), gin.Default()}
}

func (s *server) Start() {
	s.initRouters()
	s.Router.Run(s.Config.Port)
}

func (s *server) initRouters() {
	s.Router.GET("/ping", ping)
	s.Router.GET("/testing/testclient", testClient)
}
