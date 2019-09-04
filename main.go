package main

import (
	"runtime"

	"crypto-user/api/middleware"
	"crypto-user/api/user_api"
	"crypto-user/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"x-url-path", "content-type", "Authorization"}
	config.AllowMethods = []string{"POST", "OPTIONS", "GET"}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	userGroup := r.Group("/api/user")
	userGroup.POST("/create", user_api.CreateUserHandler)
	// userGroup.POST("/delete", user_api.DeleteUserHandler)
	userGroup.POST("/setkey", user_api.SetKeyUserHandler)
	userGroup.POST("/login", user_api.LoginUserHandler)
	userGroup.POST("/sms", user_api.SMSHandler)
	userGroup.POST("/setpush", user_api.SetPushURLUserHandler)
	userGroup.Use(middleware.JwtMiddleware().MiddlewareFunc())
	{
		userGroup.POST("/refresh", user_api.RefreshTokenHandler)
	}
	port, _ := utils.GetConfig().Get("user.port")
	mode, _ := utils.GetConfig().Get("gin.mode")
	if mode == "" {
		mode = "debug"
	}
	if port == "" {
		port = "5001"
	}
	gin.SetMode(mode)
	r.Run(":" + port)
}
