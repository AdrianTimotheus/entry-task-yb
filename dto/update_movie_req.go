package dto

type UpdateMovieReq struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
}
