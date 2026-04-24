package film_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alinasheleg/kinotover-go/internal/core/domain"
	core_http_response "github.com/alinasheleg/kinotover-go/internal/core/http/response"
)

func (h *filmHandler) GetFilms(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, err := strconv.Atoi(q.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(q.Get("size"))
	if err != nil || size < 1 {
		size = 10
	}
	country, err := strconv.Atoi(q.Get("country"))
	if err != nil {
		country = 0
	}
	category, err := strconv.Atoi(q.Get("category"))
	if err != nil {
		category = 0
	}
	filter := domain.Filter{
		Page:       page,
		Size:       size,
		CountryID:  country,
		CategoryID: category,
		Search:     q.Get("search"),
		SortBy:     q.Get("sortBy"),
		SortDir:    q.Get("sortDir"),
	}
	films, count, err := h.filmService.GetFilms(filter)
	if err != nil {
		core_http_response.WriteError(w, http.StatusInternalServerError, fmt.Sprintf("not recieved films: %v", err))
		return
	}
	core_http_response.WriteJSON(w, http.StatusOK, ResponseDTO{
		Page:  filter.Page,
		Size:  filter.Size,
		Total: count,
		Films: getFilmDTOsFromFilmDomains(films),
	})
}
	

