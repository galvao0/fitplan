package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.Use(cors.Default())

	InitRoutes(r)

	r.Run("0.0.0.0:8080")
}
