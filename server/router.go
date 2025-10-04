package server

import (
	"awesomeProject/handlers"
	"awesomeProject/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	ActorService services.ActorService
	MovieService services.MovieService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()

	//prevent cors error
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin, Content-Length, Content-Type, Authorization"},
	}))

	h := handlers.New(&handlers.HandlerConfig{
		ActorService: c.ActorService,
		MovieService: c.MovieService,
	})

	movieRoutes := router.Group("/movies")
	{
		movieRoutes.POST("", h.CreateMovie)
		movieRoutes.GET("", h.GetMovies)
		movieRoutes.GET("/:id", h.GetMovie)
		movieRoutes.PUT("", h.UpdateMovie)
		movieRoutes.DELETE("", h.DeleteMovie)
	}

	actorRoutes := router.Group("/actors")
	{
		actorRoutes.POST("", h.CreateActor)
		actorRoutes.GET("", h.GetActors)
		actorRoutes.GET("/:id", h.GetActor)
		actorRoutes.PUT("", h.UpdateActor)
		actorRoutes.DELETE("", h.DeleteActor)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return router
}
