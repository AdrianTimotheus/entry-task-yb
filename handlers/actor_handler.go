package handlers

import (
	"awesomeProject/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateActor(c *gin.Context) {
	var actorReq *dto.CreateActorReq
	if err := c.ShouldBindJSON(&actorReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please provide actor name"})
		return
	}

	actor, err := h.actorService.CreateActor(actorReq)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": actor})
}

func (h *Handler) GetActor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}

	actor, err := h.actorService.GetActor(id)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": actor})
}

func (h *Handler) GetActors(c *gin.Context) {
	actors, err := h.actorService.GetActors()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": actors})
}
func (h *Handler) DeleteActor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}

	actor, err := h.actorService.DeleteActor(id)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": actor})
}

func (h *Handler) UpdateActor(c *gin.Context) {
	var actorReq *dto.UpdateActorReq
	if err := c.ShouldBindJSON(&actorReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please provide id and updated actor name"})
	}

	actor, err := h.actorService.UpdateActor(actorReq)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": actor})
}
