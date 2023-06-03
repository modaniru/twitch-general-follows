package internal

import (
	"net/http"

	"github.com/modaniru/twitch-general-follows/src/twitch"
)

type Service struct {
	queries twitch.Queries
}

type response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Object     interface{} `json:"data"`
}

func newResponse(statusCode int, message string, object interface{}) *response {
	return &response{
		StatusCode: statusCode,
		Message:    message,
		Object:     object,
	}
}

func NewService(queries twitch.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) GetGeneralFollows(nicknames []string) *response {
	usersInfo, err := s.queries.GetUsersInfo(nicknames, "login")
	if err != nil {
		return newResponse(http.StatusBadRequest, err.Error(), nil)
	}
	generalFollows := make(map[string]bool)
	channel := make(chan *[]twitch.FollowInfo)
	go s.queries.GetFollows((*usersInfo)[0].Id, channel)

	for i := 1; i < len(*usersInfo); i++ {
		go s.queries.GetFollows((*usersInfo)[i].Id, channel)
	}
	followList := <-channel
	if followList == nil {
		return newResponse(http.StatusBadRequest, err.Error(), nil)
	}
	for _, v := range *followList {
		generalFollows[v.ToId] = true
	}
	for i := 1; i < len(*usersInfo); i++ {
		newGeneralFollows := make(map[string]bool)
		followList = <-channel
		if followList == nil {
			return newResponse(http.StatusBadRequest, err.Error(), nil)
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
		return newResponse(http.StatusBadRequest, err.Error(), nil)
	}
	return newResponse(200, "", response)
}

func (s *Service) Ping() *response {
	return newResponse(200, "ping", nil)
}
