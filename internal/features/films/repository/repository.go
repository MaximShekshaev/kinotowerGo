package film_repository

import (
	"github.com/alinasheleg/kinotover-go/internal/core/domain"
	"github.com/jmoiron/sqlx"
)

type FilmRepository interface {
	GetFilms() ([]domain.Film, error)
}



type filmRepository struct {
	DB *sqlx.DB
}

func NewFilmRepository(db *sqlx.DB) *filmRepository {
	return &filmRepository{DB: db}
}
