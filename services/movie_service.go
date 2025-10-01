package services

import (
	"awesomeProject/dto"
	"awesomeProject/models"
	"awesomeProject/repositories"
	"errors"
)

type MovieService interface {
	CreateMovie(movie *dto.CreateMovieReq) (*models.Movie, error)
	GetMovie(id int) (*models.Movie, error)
	GetMovies() ([]*models.Movie, error)
	UpdateMovie(movie *dto.UpdateMovieReq) (*models.Movie, error)
	DeleteMovie(movie *dto.DeleteMovieReq) (*models.Movie, error)
}

type movieService struct {
	movieRepository repositories.MovieRepository
}

type MovieServiceConfig struct {
	MovieRepository repositories.MovieRepository
}

func NewMovieService(config *MovieServiceConfig) MovieService {
	return &movieService{
		movieRepository: config.MovieRepository,
	}
}

func (s *movieService) CreateMovie(movie *dto.CreateMovieReq) (*models.Movie, error) {
	newMovie := models.Movie{
		Title:       movie.Title,
		Description: movie.Description,
		Year:        movie.Year,
		Genre:       movie.Genre,
	}

	createdMovie, err := s.movieRepository.CreateMovie(&newMovie)

	if err != nil {
		return nil, err
	}

	return createdMovie, nil
}

func (s *movieService) DeleteMovie(movie *dto.DeleteMovieReq) (*models.Movie, error) {
	if movie.ID == 0 && movie.Title == "" {
		return nil, errors.New("must provide a movie title or id")
	}

	toBeDeletedMovie := models.Movie{
		ID:    movie.ID,
		Title: movie.Title,
	}

	deletedMovie, err := s.movieRepository.DeleteMovie(&toBeDeletedMovie)

	if err != nil {
		return nil, err
	}

	return deletedMovie, nil
}

func (s *movieService) GetMovie(id int) (*models.Movie, error) {
	movie, err := s.movieRepository.GetMovie(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (s *movieService) GetMovies() ([]*models.Movie, error) {
	movies, err := s.movieRepository.GetMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *movieService) UpdateMovie(movie *dto.UpdateMovieReq) (*models.Movie, error) {
	updatedMovie := models.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Year:        movie.Year,
		Genre:       movie.Genre,
	}

	afterUpdatedMovie, err := s.movieRepository.UpdateMovie(&updatedMovie)
	if err != nil {
		return nil, err
	}

	return afterUpdatedMovie, nil
}
