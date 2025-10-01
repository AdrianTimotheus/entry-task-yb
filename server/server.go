package server

import (
	"awesomeProject/db"
	"awesomeProject/repositories"
	"awesomeProject/services"
	"log"
)

func Init() {
	actorRepository := repositories.NewActorRepository(&repositories.ActorRepositoryConfig{DB: db.Get()})
	actorService := services.NewActorService(&services.ActorServiceConfig{
		ActorRepository: actorRepository,
	})

	movieRepository := repositories.NewMovieRepository(&repositories.MovieRepositoryConfig{DB: db.Get()})
	movieService := services.NewMovieService(&services.MovieServiceConfig{
		MovieRepository: movieRepository,
	})

	router := NewRouter(&RouterConfig{
		ActorService: actorService,
		MovieService: movieService,
	})
	err := router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
