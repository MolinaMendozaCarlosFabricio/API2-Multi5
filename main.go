package main

import (
	"net/http"

	notifications_routes "api2-multi.com/a/src/Notifications/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	notifications_routes.NotifiactionsRoutes(r)

	handler := c.Handler(r)

	http.ListenAndServe(":9080", handler)
}