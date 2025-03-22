package main

import (
	"LibManMicroServ/auth"
	_ "LibManMicroServ/docs/auth"
	"LibManMicroServ/events"
	"LibManMicroServ/middleware"

	// "time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			AuthServer
//	@version		1.0
//	@description	Authentication Server
//	@BasePath	/auth
func AuthServer(eventBus *events.EventBus) *http.Server {

	PORT := viper.GetString("PORT.AUTH")

	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(middleware.RateLimiterMiddleware(limiter.Rate{
	// 	Period: 1 * time.Minute,
	// 	Limit:  5,
	// }))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("auth")))
	auth.Router(eventBus, r)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server

}
