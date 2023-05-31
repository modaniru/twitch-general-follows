package twitch

import (
	"encoding/json"
	"fmt"
	"io"
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
}

func NewQueries(cfgPath string) *Queries {
	
	queries := Queries{}
	toml.DecodeFile(cfgPath, &queries)
	return &queries
}

func (q *Queries) GetOauthToken() (*OauthToken, error) {
	timeout := 10 * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	uri := fmt.Sprintf("%s?client_id=%s&client_secret=%s&grant_type=client_credentials", q.GetTokenURI, q.ClientId, q.ClientSecret)
	fmt.Println(uri)
	res, err := client.Post(uri, nil, nil)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	var response OauthToken
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}
