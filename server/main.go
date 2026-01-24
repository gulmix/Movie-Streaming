package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gulmix/Movie-Streaming/server/routes"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("failed to start server", err)
	}
}
