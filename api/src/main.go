package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"routes"
)

func main() {
	r := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:4000"}
	r.Use(cors.New(config))
	
	r.GET("/", routes.Index)
	r.POST("/login", routes.Login)

	r.Run()
}