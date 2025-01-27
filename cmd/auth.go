package main

import (
	"LibManMicroServ/auth"
	"LibManMicroServ/events"
	"LibManMicroServ/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthServer(eventBus *events.EventBus) *http.Server {

	PORT := viper.GetString("PORT.AUTH")

	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	auth.Router(eventBus, r)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server

}
