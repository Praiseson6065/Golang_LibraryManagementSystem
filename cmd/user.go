package main

import (
	"LibManMicroServ/lending"
	"LibManMicroServ/middleware"
	"LibManMicroServ/reviews"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func UserReviewServer() *http.Server {
	PORT := viper.GetString("PORT.USER.REVIEWS")
	r := gin.New()
	r.Use(middleware.Authenicator())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	reviews.UserRouter(r)
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server
}

func UserLendingServer() *http.Server {
	PORT := viper.GetString("PORT.USER.LENDING")
	r := gin.New()
	r.Use(middleware.Authenicator())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	lending.UserRouter(r)
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server
}
