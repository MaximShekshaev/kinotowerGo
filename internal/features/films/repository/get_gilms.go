package film_repository

import (
	"fmt"
	"strings"

	"github.com/alinasheleg/kinotover-go/internal/core/domain"
	"github.com/jmoiron/sqlx"
)

func (r *filmRepository) GetFilms(filter domain.Filter) ([]domain.Film, int, error) {
	query, args := buildFilmsQuery(filter)

	rows := []Film{}
	if err := r.DB.Select(&rows, query, args...); err != nil {
		return nil, 0, fmt.Errorf("not received films: %w", err)
	}

	// collect film IDs
	filmIDs := make([]int, len(rows))
	for i, row := range rows {
		filmIDs[i] = row.ID

		rows[i].Country = Country{
			ID:   row.CountryID,
			Name: row.CountryName,
		}
	}

	// categories
	categoriesMap, err := r.GetCategories(filmIDs)
	if err != nil {
		return nil, 0, fmt.Errorf("not received categories for films: %w", err)
	}

	for i, row := range rows {
		rows[i].Categories = categoriesMap[row.ID]
	}

	// count
	count, err := r.CountFilms(filter)
	if err != nil {
		return nil, 0, fmt.Errorf("not received count of films: %w", err)
	}

	return getFilmDomainsFromFilmModels(rows), count, nil
}

func (r *filmRepository) CountFilms(filter domain.Filter) (int, error) {
	args := []interface{}{}
	where := []string{"f.deleted_at IS NULL"}

	if filter.CountryID != 0 {
		where = append(where, fmt.Sprintf("f.country_id = $%d", len(args)+1))
		args = append(args, filter.CountryID)
	}

	if filter.Search != "" {
		where = append(where, fmt.Sprintf("f.name ILIKE $%d", len(args)+1))
		args = append(args, "%"+filter.Search+"%")
	}

	if filter.CategoryID != 0 {
		where = append(where, fmt.Sprintf(`
			EXISTS (
				SELECT 1
				FROM categories_films cf
				WHERE cf.film_id = f.id AND cf.category_id = $%d
			)
		`, len(args)+1))
		args = append(args, filter.CategoryID)
	}

	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM films f
		WHERE %s
	`, strings.Join(where, " AND "))

	var count int
	if err := r.DB.Get(&count, query, args...); err != nil {
		return 0, fmt.Errorf("not received count of films: %w", err)
	}

	return count, nil
}

func (r *filmRepository) GetCategories(filmIDs []int) (map[int][]Category, error) {
	if len(filmIDs) == 0 {
		return map[int][]Category{}, nil
	}

	query, args, err := sqlx.In(`
		SELECT
			cf.film_id,
			c.id,
			c.name
		FROM categories_films cf
		JOIN categories c ON cf.category_id = c.id
		WHERE cf.film_id IN (?)
	`, filmIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to build categories query: %w", err)
	}

	query = r.DB.Rebind(query)

	type categoryRow struct {
		FilmID int `db:"film_id"`
		Category
	}

	rows := []categoryRow{}
	if err := r.DB.Select(&rows, query, args...); err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	categoryMap := make(map[int][]Category)
	for _, row := range rows {
		categoryMap[row.FilmID] = append(categoryMap[row.FilmID], row.Category)
	}

	return categoryMap, nil
}

func buildFilmsQuery(filter domain.Filter) (string, []interface{}) {
	args := []interface{}{}
	where := []string{"f.deleted_at IS NULL"}

	if filter.CountryID != 0 {
		where = append(where, fmt.Sprintf("f.country_id = $%d", len(args)+1))
		args = append(args, filter.CountryID)
	}

	if filter.Search != "" {
		where = append(where, fmt.Sprintf("f.name ILIKE $%d", len(args)+1))
		args = append(args, "%"+filter.Search+"%")
	}

	if filter.CategoryID != 0 {
		where = append(where, fmt.Sprintf(`
			EXISTS (
				SELECT 1
				FROM categories_films cf
				WHERE cf.film_id = f.id AND cf.category_id = $%d
			)
		`, len(args)+1))
		args = append(args, filter.CategoryID)
	}

	// LIMIT
	limitIdx := len(args) + 1
	args = append(args, filter.Limit())
	limit := fmt.Sprintf("LIMIT $%d", limitIdx)

	// OFFSET
	offsetIdx := len(args) + 1
	args = append(args, filter.Offset())
	offset := fmt.Sprintf("OFFSET $%d", offsetIdx)

	// ORDER BY
	orderBy := buildOrderBy(filter)

	query := fmt.Sprintf(`
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
			c.name AS country_name,
			COALESCE(AVG(r.ball), 0) AS rating_avg,
			COUNT(rv.id) AS review_count
		FROM films f
		JOIN countries c ON f.country_id = c.id
		LEFT JOIN ratings r ON f.id = r.film_id
		LEFT JOIN reviews rv ON f.id = rv.film_id AND rv.deleted_at IS NULL
		WHERE %s
		GROUP BY f.id, c.id
		%s
		%s
		%s
	`,
		strings.Join(where, " AND "),
		orderBy,
		limit,
		offset,
	)

	return query, args
}

func buildOrderBy(filter domain.Filter) string {
	allowedSortBy := map[string]string{
		"name":   "f.name",
		"year":   "f.year_of_issue",
		"rating": "rating_avg",
	}

	sortBy, ok := allowedSortBy[filter.SortBy]
	if !ok {
		sortBy = "f.name"
	}

	sortDir := "ASC"
	if strings.ToLower(filter.SortDir) == "desc" {
		sortDir = "DESC"
	}

	return fmt.Sprintf("ORDER BY %s %s", sortBy, sortDir)
}