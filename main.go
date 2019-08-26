package main

import (
	"runtime"

	"./api/user_api"
	"./utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"x-url-path", "content-type"}
	config.AllowMethods = []string{"POST", "OPTIONS", "GET"}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	userGroup := r.Group("/api/user")
	userGroup.POST("/create", user_api.CreateUserHandler)
	userGroup.POST("/delete", user_api.DeleteUserHandler)

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
