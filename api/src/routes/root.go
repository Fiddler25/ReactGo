package routes

import "github.com/gin-gonic/gin"

func Route(c *gin.Context) {
	c.JSON(200, gin.H{"title": "Hello World!"})
}
