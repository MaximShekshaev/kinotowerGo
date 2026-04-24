package film_service

import (
	"strings"

	"github.com/alinasheleg/kinotover-go/internal/core/domain"
)

func (s *filmService) GetFilms(filter domain.Filter) ([]domain.Film, int, error) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Size <= 0 {
		filter.Size = 3
	}

	filter.SortBy = strings.ToLower(filter.SortBy)
	filter.SortDir = strings.ToLower(filter.SortDir)

	return s.filmRepository.GetFilms(filter)
}