package core_server

import (
	"net/http"

	core_router "github.com/alinasheleg/kinotover-go/internal/core/router"
	film_handler "github.com/alinasheleg/kinotover-go/internal/features/films/handler"
	film_repository "github.com/alinasheleg/kinotover-go/internal/features/films/repository"
	film_service "github.com/alinasheleg/kinotover-go/internal/features/films/service"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	http.Server
}

func NewServer(db *sqlx.DB) *Server {
	cfg := NewConfigMust()

	filmRepository := film_repository.NewFilmRepository(db)
	filmService := film_service.NewFilmService(filmRepository)
	filmHandler := film_handler.NewFilmHandler(filmService)


	router := core_router.NewRouter(filmHandler)
	return &Server{
		Server: http.Server{
			Addr:    cfg.Addr,
			Handler: router.RegisterRoute(),
		},
	}

}
