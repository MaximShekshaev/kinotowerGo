package domain

type Film struct {
	ID 		int    
	Name 	string 
	Duration int 
	YearOfissue int
	Age 	int
	LinkImg 	string
	LinkKinopoisk string
	LinlVideo string
	CreatedAt string
	Country Country
	Categories []Category
	RatingAvg float64
	ReviewCount int
}
