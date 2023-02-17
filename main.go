package main

import (
	"golang-example/config"
	"golang-example/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	router := gin.Default()

	routes.RegisterUserRoutes(router)

	// Run migrations
	// migrations.AutoMigrate()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}
