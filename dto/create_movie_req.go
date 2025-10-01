package dto

type CreateMovieReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
}
