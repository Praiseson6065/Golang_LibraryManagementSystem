package main

import (
	"LibManMicroServ/middleware"
	"LibManMicroServ/payments"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func PaymentsServer() *http.Server {
	PORT := viper.GetString("PORT.PAYMENTS")

	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(middleware.Authenicator())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	payments.Router(r)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server

}
