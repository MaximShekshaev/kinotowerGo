package film_repository

import (
	"fmt"

	"github.com/alinasheleg/kinotover-go/internal/core/domain"
)

func (r *filmRepository) GetFilmByID(id int) (domain.Film, error) {
	query := `
		SELECT
			f.id,
			f.name,
			f.duration,
			f.year_of_issue,
			f.age,
			f.link_img,
			f.link_kinopoisk,
			f.link_video,
			f.created_at,
			f.country_id,
			c.name AS country_name
		FROM films f
		JOIN countries c ON f.country_id = c.id
		WHERE f.id = $1 AND f.deleted_at IS NULL
	`

	var row Film
	if err := r.DB.Get(&row, query, id); err != nil {
		return domain.Film{}, fmt.Errorf("film not found: %w", err)
	}

	// country
	row.Country = Country{
		ID:   row.CountryID,
		Name: row.CountryName,
	}

	// categories
	categoriesMap, err := r.GetCategories([]int{id})
	if err != nil {
		return domain.Film{}, fmt.Errorf("failed to get categories: %w", err)
	}

	row.Categories = categoriesMap[id]

	// domain mapping
	return getFilmDomainFromFilmModel(row), nil
}