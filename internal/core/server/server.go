package core_server

import (
	"net/http"

	core_router "github.com/alinasheleg/kinotover-go/internal/core/router"
)

type Server struct {

	http.Server
}
func NewServer() *Server {
	router := core_router.NewRouter()


	return &Server{
	Server: http.Server{
		Addr: ":8080",
		Handler: router.RegisterRoutes(),
	},
	}

}