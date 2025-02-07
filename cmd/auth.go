package main

import (
	"LibManMicroServ/auth"
	_ "LibManMicroServ/docs/auth"
	"LibManMicroServ/events"
	"LibManMicroServ/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			AuthServer
//	@version		1.0
//	@description	Authentication Server
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/
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
