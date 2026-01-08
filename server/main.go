package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gulmix/Movie-Streaming/server/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/movie", controllers.AddMovie())

	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("failed to start server", err)
	}
}
