package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gulmix/Movie-Streaming/server/routes"
)

func main() {
	router := gin.Default()

	config := cors.Config{}
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("failed to start server", err)
	}
}
