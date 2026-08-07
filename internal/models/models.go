package models

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Movie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"releaseYear"`
	Duration    int    `json:"duration"`
}

type Actor struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
}
