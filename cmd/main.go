package main

import (
	_ "LibManMicroServ/config"

	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	gin.SetMode(gin.ReleaseMode)	

	g.Go(func() error {
		return AuthServer().ListenAndServe()
	})

	g.Go(func() error {
		return PaymentsServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
