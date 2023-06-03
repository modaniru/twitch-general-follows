package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/twitch-general-follows/src/twitch"
)

type Service struct {
	queries twitch.Queries
}

func NewService(queries twitch.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) GetGeneralFollows(c *gin.Context) {
	nicknames := c.QueryArray("login")
	usersInfo, err := s.queries.GetUsersInfo(nicknames, "login")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	generalFollows := make(map[string]bool)
	channel := make(chan *[]twitch.FollowInfo)
	go s.queries.GetFollows((*usersInfo)[0].Id, channel)

	for i := 1; i < len(*usersInfo); i++ {
		go s.queries.GetFollows((*usersInfo)[i].Id, channel)
	}
	followList := <-channel
	if followList == nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	for _, v := range *followList {
		generalFollows[v.ToId] = true
	}
	for i := 1; i < len(*usersInfo); i++ {
		newGeneralFollows := make(map[string]bool)
		followList = <-channel
		if followList == nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		for _, v := range *followList {
			if generalFollows[v.ToId] {
				newGeneralFollows[v.ToId] = true
			}
		}
		generalFollows = newGeneralFollows
	}
	idsList := make([]string, len(generalFollows))
	i := 0
	for k := range generalFollows {
		idsList[i] = k
		i++
	}
	response, err := s.queries.GetUsersInfo(idsList, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

func (s *Service) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ping"})
}
