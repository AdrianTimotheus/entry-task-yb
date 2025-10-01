package handlers

import "awesomeProject/services"

type Handler struct {
	actorService services.ActorService
	movieService services.MovieService
}

type HandlerConfig struct {
	ActorService services.ActorService
	MovieService services.MovieService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		actorService: c.ActorService,
		movieService: c.MovieService,
	}
}
