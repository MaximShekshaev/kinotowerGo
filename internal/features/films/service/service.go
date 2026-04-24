package film_service

import (
	"github.com/MaximShekshaev/kinotowerGo/internal/core/domain"
	film_repository "github.com/MaximShekshaev/kinotowerGo/internal/features/films/repository"
)

type FilmService interface {
	GetFilms(filter domain.Filter) ([]domain.Film, int, error)
}

type filmService struct {
	filmRepository film_repository.FilmRepository
}

func NewFilmService(filmRepo film_repository.FilmRepository) *filmService {
	return &filmService{
		filmRepository: filmRepo,
	}
}