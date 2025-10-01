package repositories

import (
	"awesomeProject/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movie *models.Movie) (*models.Movie, error)
	GetMovie(id int) (*models.Movie, error)
	GetMovies() ([]*models.Movie, error)
	UpdateMovie(movie *models.Movie) (*models.Movie, error)
	DeleteMovie(movie *models.Movie) (*models.Movie, error)
}

type movieRepository struct {
	db *gorm.DB
}

type MovieRepositoryConfig struct {
	DB *gorm.DB
}

func NewMovieRepository(config *MovieRepositoryConfig) MovieRepository {
	return &movieRepository{db: config.DB}
}

func (r *movieRepository) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	result := r.db.Create(movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}

func (r *movieRepository) GetMovie(id int) (*models.Movie, error) {
	var movie models.Movie
	result := r.db.Where("id = ?", id).First(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func (r *movieRepository) GetMovies() ([]*models.Movie, error) {
	var movies []*models.Movie
	result := r.db.Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	//add comment
	return movies, nil
}

func (r *movieRepository) UpdateMovie(movie *models.Movie) (*models.Movie, error) {
	result := r.db.Save(movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}

func (r *movieRepository) DeleteMovie(movie *models.Movie) (*models.Movie, error) {
	var deletedMovie models.Movie
	result := r.db.Where("id = ?", movie.ID).First(&deletedMovie)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Where("id = ?", movie.ID).Delete(&movie)
	if result.Error != nil {
		return nil, result.Error
	}

	return &deletedMovie, nil
}
