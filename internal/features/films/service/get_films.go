package film_service

import "github.com/alinasheleg/kinotover-go/internal/core/domain"

func (s *filmService) GetFilms() ([]domain.Film, error) {
	return s.filmRepository.GetFilms()
}