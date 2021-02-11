package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"routes"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4000"}
	router.Use(cors.New(config))

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
