package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gulmix/Movie-Streaming/server/controllers"
	"github.com/gulmix/Movie-Streaming/server/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser())
	router.POST("/login", controllers.LoginUser())
	router.GET("/movies", controllers.GetMovies())

	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/movie", controllers.AddMovie())
	router.GET("/recommended", controllers.GetRecommendedMovies())
	router.PATCH("/review/:imdb_id", controllers.AdminReviewUpdate())
}
