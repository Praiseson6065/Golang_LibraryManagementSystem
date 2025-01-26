package main

import (
	"LibManMicroServ/books"
	"LibManMicroServ/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AdminBooksServer() *http.Server {
	PORT := viper.GetString("PORT.ADMIN.BOOKS")
	r := gin.New()
	r.Use(middleware.Authenicator())
	r.Use(middleware.EnsureAdmin())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	books.AdminRouter(r)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server

}
