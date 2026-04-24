package domain

import "time"

type Film struct {
	ID            int
	Name          string
	Duration      int
	YearOfIssue   int
	Age           int
	LinkImg       *string
	LinkKinopoisk *string
	LinkVideo     string
	CreatedAt     time.Time
	Country       Country
	Categories    []Category
	RatingAvg     float64
	ReviewCount   int
}
