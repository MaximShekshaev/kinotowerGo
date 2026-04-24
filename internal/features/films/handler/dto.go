package film_handler

import (
	"time"

	"github.com/alinasheleg/kinotover-go/internal/core/domain"
)

type FilmDTO struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Duration      int           `json:"duration"`
	YearOfIssue   int           `json:"year_of_issue"`
	Age           int           `json:"age"`
	LinkImg       *string       `json:"link_img,omitempty"`
	LinkKinopoisk *string       `json:"link_kinopoisk,omitempty"`
	LinkVideo     string        `json:"link_video"`
	CreatedAt     time.Time     `json:"created_at"`
	Country       CountryDTO    `json:"country"`
	Categories    []CategoryDTO `json:"categories"`
	RatingAvg     float64       `json:"rating_avg"`
	ReviewCount   int           `json:"review_count"`
}

type CountryDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseDTO struct {
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Total int       `json:"total"`
	Films []FilmDTO `json:"films"`
}

// --- MAPPERS ---

func NewFilmDTO(film domain.Film) FilmDTO {
	return FilmDTO{
		ID:            film.ID,
		Name:          film.Name,
		Duration:      film.Duration,
		YearOfIssue:   film.YearOfIssue, 
		Age:           film.Age,
		LinkImg:       film.LinkImg,
		LinkKinopoisk: film.LinkKinopoisk,
		LinkVideo:     film.LinkVideo,
		CreatedAt:     film.CreatedAt,
		Country:       NewCountryDTO(film.Country),
		Categories:    NewCategoryDTOs(film.Categories),
		RatingAvg:     film.RatingAvg,
		ReviewCount:   film.ReviewCount,
	}
}

func NewFilmDTOs(films []domain.Film) []FilmDTO {
	if len(films) == 0 {
		return []FilmDTO{}
	}

	result := make([]FilmDTO, len(films))
	for i, film := range films {
		result[i] = NewFilmDTO(film)
	}
	return result
}

func NewCountryDTO(country domain.Country) CountryDTO {
	return CountryDTO{
		ID:   country.ID,
		Name: country.Name,
	}
}

func NewCategoryDTO(category domain.Category) CategoryDTO {
	return CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}

func NewCategoryDTOs(categories []domain.Category) []CategoryDTO {
	if len(categories) == 0 {
		return []CategoryDTO{}
	}

	result := make([]CategoryDTO, len(categories))
	for i, category := range categories {
		result[i] = NewCategoryDTO(category)
	}
	return result
}




func getFilmDTOsFromFilmDomains(films []domain.Film) []FilmDTO{
	
 
	filmDTOs := make([]FilmDTO, len(films))
	for i, film := range films {
		filmDTOs[i] = NewFilmDTO(film)
	}
	return filmDTOs
}