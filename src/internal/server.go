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
	s.Router.GET("/get/", s.getGeneralFollows)
}

func (s *server) ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ping"})
}

func (s *server) getGeneralFollows(c *gin.Context) {
	nicknames := c.QueryArray("login")
	response, err := s.queries.GetUsersInfo(nicknames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	result := make(map[string]string)
	resp, err := s.queries.GetFollows(response.Data[0].Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	for _, v := range *resp{
		result[v.ToName] = v.ToName
	}

	for i := 1; i < len(response.Data); i++{
		newMap := make(map[string]string)
		resp, err := s.queries.GetFollows(response.Data[i].Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		for _, v := range *resp{
			_, ok := result[v.ToName]
			if ok {
				newMap[v.ToName] = v.ToName
			}
		}
		result = newMap

	}
	c.JSON(http.StatusOK, result)
}

func (s *server) GetFollows(c *gin.Context){
	response, err := s.queries.GetFollows("171985899")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
