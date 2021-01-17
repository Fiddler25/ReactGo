package main

import (
	"github.com/gin-gonic/gin"
	"routes"
)

func main() {
	r := gin.Default()
	r.GET("/", routes.Index)

	r.Run()
}