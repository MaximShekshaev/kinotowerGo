package film_repository

import (
	"time"

	"github.com/alinasheleg/kinotover-go/internal/core/domain"
)

type Film struct {
	ID            int     `db:"id"`
	Name          string  `db:"name"`
	Duration      int     `db:"duration"`
	YearOfIssue   int     `db:"year_of_issue"`
	Age           int     `db:"age"`
	LinkImg       *string `db:"link_img"`
	LinkKinopoisk *string `db:"link_kinopoisk"`
	LinkVideo     string  `db:"link_video"`
	CreatedAt     string  `db:"created_at"`
	CountryID     int     `db:"country_id"`
	CountryName   string  `db:"country_name"`
	RatingAvg     float64 `db:"rating_avg"`
	ReviewCount   int     `db:"review_count"`

	Country    Country
	Categories []Category
}
type Country struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func getFilmDomainFromFilmModel(filmModel Film) domain.Film {
	categories := make([]domain.Category, len(filmModel.Categories))
	for i, category := range filmModel.Categories {
		categories[i] = domain.Category{
			ID:   category.ID,
			Name: category.Name,
		}
	}
	createdAt , _ := time.Parse(time.RFC3339, filmModel.CreatedAt)
	return domain.Film{
		ID:            filmModel.ID,
		Name:          filmModel.Name,
		Duration:      filmModel.Duration,
		YearOfIssue:   filmModel.YearOfIssue,
		Age:           filmModel.Age,
		LinkImg:       filmModel.LinkImg,
		LinkKinopoisk: filmModel.LinkKinopoisk,
		LinkVideo:     filmModel.LinkVideo,
		CreatedAt:     createdAt,
		Country: domain.Country{
			ID:   filmModel.CountryID,
			Name: filmModel.CountryName,
		},
		Categories: categories,
	}
}

func getFilmDomainsFromFilmModels(filmModels []Film) []domain.Film {
	films := make([]domain.Film, len(filmModels))
	for i, filmModel := range filmModels {
		films[i] = getFilmDomainFromFilmModel(filmModel)
	}
	return films
}