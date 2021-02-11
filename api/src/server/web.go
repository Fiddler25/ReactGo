package server

import (
	"github.com/gin-gonic/gin"
	"routes"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	endpoints(router)

	return router
}

func endpoints(r *gin.Engine) {
	r.GET("/", routes.Route)
}
