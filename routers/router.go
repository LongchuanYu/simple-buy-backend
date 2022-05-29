package routers

import (
	cors "simplebuy/middlewares"
	login "simplebuy/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())
	gin.SetMode("debug")
	api := r.Group("/api")
	{
		api.GET("/login", login.Login)
	}

	return r
}