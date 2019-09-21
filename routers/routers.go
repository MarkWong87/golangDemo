package routers

import (
	v1 "golang/api/v1"
	v2 "golang/api/v2"
	"golang/middleware"
	"golang/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("api/v1")
	{
		apiV1.Use(middleware.Auth())

		apiV1.POST("/login", v1.Login)
		apiV1.POST("/logout", v1.Logout)
		apiV1.GET("/code", v1.GetCode)
	}

	apiV2 := r.Group("api/v2")
	{
		apiV2.Use(middleware.Token())

		apiV2.Use(middleware.Auth())
		apiV2.POST("/login", v2.Login)
		apiV2.POST("/logout", v2.Logout)
		apiV2.GET("/code", v2.GetCode)
	}

	return r
}
