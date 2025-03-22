package main

import (
	"LibManMicroServ/cart"
	"LibManMicroServ/events"
	"LibManMicroServ/lending"
	"LibManMicroServ/middleware"
	"LibManMicroServ/reviews"
	"net/http"

	_ "LibManMicroServ/docs/usercart"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func newRouter(swaggerInstance string) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Authenicator())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName(swaggerInstance)))
	return r
}

func UserReviewServer() *http.Server {
	PORT := viper.GetString("PORT.USER.REVIEWS")
	r := newRouter("user")
	reviews.UserRouter(r)
	return &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
}

func UserLendingServer() *http.Server {
	PORT := viper.GetString("PORT.USER.LENDING")
	r := newRouter("user")
	lending.UserRouter(r)
	return &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
}

//	@title			UserCartServer
//	@version		1.0
//	@description	User Cart Server
//	@BasePath	/cart
//	@securityDefinitions.apikey BearerAuth
// 	@in header
// 	@name Authorization
func UserCartServer(eventBus *events.EventBus) *http.Server {
	PORT := viper.GetString("PORT.USER.CART")
	r := newRouter("usercart")
	cart.Router(eventBus, r)
	return &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
}
