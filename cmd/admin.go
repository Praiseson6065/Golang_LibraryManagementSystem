package main

import (
	"LibManMicroServ/books"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AdminServer() *http.Server {
	PORT := viper.GetString("PORT.ADMIN")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use()
	books.AdminRouter(r)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server

}
