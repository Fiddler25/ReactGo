package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
	"routes"
)

func main() {
	r := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:4000"}
	r.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))
 
    r.GET("/hello", func(c *gin.Context) {
        session := sessions.Default(c)
 
        if session.Get("hello") != "world" {
            session.Set("hello", "world")
            session.Save()
        }
 
        c.JSON(200, gin.H{ "hello": session.Get("hello") })
    })
	
	r.GET("/", routes.Index)
	r.POST("/login", routes.Login)

	r.Run()
}