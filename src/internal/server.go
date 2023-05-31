package internal

import "github.com/gin-gonic/gin"

type server struct{
	Config *configauration
	Router *gin.Engine
}

func NewServer(configPath string) *server{
	return &server{NewConfiguration(configPath), gin.Default()}
}

func (s *server)Start(){
	//todo routing methods
	s.Router.Run(s.Config.Port)
}