package dto

type CreateMovieReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Year        int    `json:"year" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
}
