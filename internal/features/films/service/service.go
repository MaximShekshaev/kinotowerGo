package film_service

import (
	"github.com/alinasheleg/kinotover-go/internal/core/domain"
	film_repository "github.com/alinasheleg/kinotover-go/internal/features/films/repository"
)

type FilmService interface {
	GetFilms() ([]domain.Film, error)
}

type filmService struct {
	filmRepository film_repository.FilmRepository
}

func NewFilmService(filmRepo film_repository.FilmRepository) *filmService {
	return &filmService{
		filmRepository: filmRepo,
	}
}