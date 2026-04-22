package film_repository

import "github.com/alinasheleg/kinotover-go/internal/core/domain"

func (r *filmRepository) GetFilms() ([]domain.Film, error) {
	return []domain.Film{}, nil
}