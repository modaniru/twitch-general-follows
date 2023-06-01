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
	result := make(map[string]bool)
	channel := make(chan *[]twitch.FollowInfo)
	go s.queries.GetFollows((*response)[0].Id, channel)

	for i := 1; i < len(*response); i++ {
		go s.queries.GetFollows((*response)[i].Id, channel)
	}
	list := <- channel
	if list == nil{
		c.JSON(http.StatusInternalServerError, "")
	}
	for _, v := range *list{
		result[v.ToName] = true
	}
	for i := 1; i < len(*response); i++{
		temp := make(map[string]bool)
		list = <- channel
		if list == nil{
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		for _, v := range *list{
			if result[v.ToName]{
				temp[v.ToName] = true
			}
		}
		result = temp
	}
	c.JSON(http.StatusOK, result)
}
