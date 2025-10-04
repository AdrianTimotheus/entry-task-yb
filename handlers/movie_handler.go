package handlers

import (
	"awesomeProject/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	var movieReq *dto.CreateMovieReq
	if err := c.ShouldBindJSON(&movieReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movieService.CreateMovie(movieReq)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"movie": movie})
}

func (h *Handler) UpdateMovie(c *gin.Context) {
	var movieReq *dto.UpdateMovieReq
	if err := c.ShouldBindJSON(&movieReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	movie, err := h.movieService.UpdateMovie(movieReq)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

func (h *Handler) DeleteMovie(c *gin.Context) {
	var movieReq *dto.DeleteMovieReq
	if err := c.ShouldBindJSON(&movieReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	movie, err := h.movieService.DeleteMovie(movieReq)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

func (h *Handler) GetMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}

	movie, err := h.movieService.GetMovie(id)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}
func (h *Handler) GetMovies(c *gin.Context) {
	genre := c.Query("genre")
	movies, err := h.movieService.GetMovies(genre)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}
