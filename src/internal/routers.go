package internal

import "github.com/gin-gonic/gin"

func ping(c *gin.Context){
	c.JSON(200, gin.H{"message": "ping"})
}