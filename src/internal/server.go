package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/twitch-general-follows/src/twitch"
)

type server struct {
	Config  *configauration
	Router  *gin.Engine
	queries *twitch.Queries
}

func NewServer(configPath, twtichCfgPath string) *server {
	return &server{Config: NewConfiguration(configPath), Router: gin.Default(), queries: twitch.NewQueries(twtichCfgPath)}
}

func (s *server) Start() {
	s.initRouters()
	s.Router.Run(s.Config.Port)
}

func (s *server) initRouters() {
	s.Router.GET("/ping", s.ping)
	s.Router.GET("/oauth", s.getOauthToken)
}

func (s *server) ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ping"})
}

func (s *server) getOauthToken(c *gin.Context) {
	response, err := s.queries.GetOauthToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
