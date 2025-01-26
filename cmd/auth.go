package main

import (
	"LibManMicroServ/auth"
	"LibManMicroServ/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthServer() *http.Server {

	PORT := viper.GetString("PORT.AUTH")
	fmt.Println("Auth Server started at PORT: ", PORT)
	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	auth.Router(r)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server

}
