package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-contrib/sessions/cookie"
	"queries"
	"server"
)

func main() {
	router := gin.Default()

	queries.Query()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
 
	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{ "hello": session.Get("hello") })
	})

	server.Init()
}