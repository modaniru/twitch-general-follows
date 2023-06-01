package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gojek/heimdall/v7/httpclient"
)

type Queries struct {
	UserInfoURI          string `toml:"user_info"`
	UserGetFollowListURI string `toml:"user_get_follow_list"`
	GetTokenURI          string `toml:"get_token"`
	ValidateToken        string `toml:"validate_token"`

	ClientId     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`

	client *httpclient.Client
	token  string
}

func NewQueries(cfgPath string) *Queries {
	timeout := 10 * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	queries := Queries{client: client}
	toml.DecodeFile(cfgPath, &queries)
	token, err := queries.GetOauthToken()
	if err != nil {
		log.Fatal(err)
	}
	queries.token = token.AccessToken
	return &queries
}

func (q *Queries) GetOauthToken() (*OauthToken, error) {
	uri := fmt.Sprintf("%s?client_id=%s&client_secret=%s&grant_type=client_credentials", q.GetTokenURI, q.ClientId, q.ClientSecret)
	fmt.Println(uri)
	res, err := q.client.Post(uri, nil, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		message := fmt.Sprintf("%d status code", res.StatusCode)
		return nil, errors.New(message)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response OauthToken
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}

func (q *Queries) IsValid() (*ValidToken, error) {
	uri := q.ValidateToken
	token := "OAuth " + q.token
	header := http.Header{}
	header.Add("Authorization", token)
	res, err := q.client.Get(uri, header)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		message := fmt.Sprintf("%d status code", res.StatusCode)
		return nil, errors.New(message)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response ValidToken
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}

func (q *Queries) GetUsersInfo(nicknames []string) (*UserCollection, error) {
	uri := q.UserInfoURI
	for i, v := range nicknames {
		symb := "&"
		if i == 0 {
			symb = "?"
		}
		uri = fmt.Sprintf("%s%slogin=%s", uri, symb, v)
	}
	header := http.Header{}
	token := "Bearer " + "udbafm1cghmrgy9aw9xf707360ibwp"
	header.Add("Authorization", token)
	header.Add("Client-Id", q.ClientId)
	res, err := q.client.Get(uri, header)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		message := fmt.Sprintf("%d status code", res.StatusCode)
		return nil, errors.New(message)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response UserCollection
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}

func (q *Queries) getToken() string {
	_, err := q.IsValid()
	if err != nil {
		token, err := q.GetOauthToken()
		if err == nil {
			return ""
		}
		q.token = token.AccessToken
	}
	return q.token
}

// go routine
func (q *Queries) GetFollows(id string) (*[]FollowInfo, error) {
	uri := fmt.Sprintf("%s?from_id=%s&first=%d", q.UserGetFollowListURI, id, 100)
	response, err := q.getFollowsWithoutPagination(uri)
	if err != nil{
		return nil, err
	}
	var result []FollowInfo
	result = append(result, response.Data...)
	for response.Pagination.Cursor != ""{
		uri2 := fmt.Sprintf("%s&after=%s", uri, response.Pagination.Cursor)
		response, err = q.getFollowsWithoutPagination(uri2)
		if err != nil{
			return nil, err
		}
		result = append(result, response.Data...)
	}
	return &result, nil
}

func (q *Queries) getFollowsWithoutPagination(uri string) (*FollowsCollection, error){
	header := http.Header{}
	token := "Bearer " + q.getToken()
	header.Add("Authorization", token)
	header.Add("Client-Id", q.ClientId)
	res, err := q.client.Get(uri, header)
	if err != nil{
		return nil, err
	}
	if res.StatusCode != 200 {
		message := fmt.Sprintf("%d status code", res.StatusCode)
		return nil, errors.New(message)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response FollowsCollection
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}
