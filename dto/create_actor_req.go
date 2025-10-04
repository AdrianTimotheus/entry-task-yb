package dto

type CreateActorReq struct {
	Name string `json:"name" binding:"required"`
}
