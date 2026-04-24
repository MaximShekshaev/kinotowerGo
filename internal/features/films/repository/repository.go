package film_repository

import (
	"github.com/alinasheleg/kinotover-go/internal/core/domain"
	"github.com/jmoiron/sqlx"
)

type FilmRepository interface {
	GetFilms(filter domain.Filter) ([]domain.Film,int, error)
}



type filmRepository struct {
	DB *sqlx.DB
}

func NewFilmRepository(db *sqlx.DB) *filmRepository {
	return &filmRepository{DB: db}
}
