package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"routes"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4000"}
	router.Use(cors.New(config))
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	endpoints(router)

	return router
}

func endpoints(r *gin.Engine) {
	r.GET("/", routes.Route)

	api := r.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			v1api.POST("/login", routes.Login)
		}
	}
}
