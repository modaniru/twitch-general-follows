package internal

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gojek/heimdall/v7/httpclient"
)

type post struct{
	Id int `json:"id"`
	UserId int `json:"userId"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ping"})
}

func testClient(c *gin.Context) {
	// Create a new HTTP client with a default timeout
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	res, err := client.Get("https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	body, err := ioutil.ReadAll(res.Body)
	var posts []post
	json.Unmarshal([]byte(string(body)), &posts)
	c.JSON(200, posts)
}
