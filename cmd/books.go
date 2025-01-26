package main

import (
	"LibManMicroServ/books"
	"LibManMicroServ/middleware"
	"LibManMicroServ/reviews"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func booksServer() *http.Server {
	PORT := viper.GetString("PORT.BOOKS")
	fmt.Println("Books Server Started at PORT - " + PORT)

	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	books.Router(r)
	reviews.Router(r)
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	return server

}
