package models

import (
	"time"

	"greenlight.letsgo.com/internals/validator"
)

type Movie struct {
	ID        int       `json:"id"`        // Unique integer ID for the movie
	CreatedAt time.Time `json:"createdAt"` // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`     // Movie title
	Year      int       `json:"year"`      // Movie release year
	Runtime   int       `json:"runtime"`   // Movie runtime (in minutes)
	Genres    []string  `json:"genres"`    // Slice of genres for the movie (romance, comedy, etc.)
	Version   int       `json:"version"`   // The version number starts at 1 and will be incremented each

}

type CreateMovieInput struct {
	Title   string   `json:"title"`   // Movie title
	Year    int      `json:"year"`    // Movie release year
	Runtime int      `json:"runtime"` // Movie runtime (in minutes)
	Genres  []string `json:"genres"`  // Slice of genres for the movi
}

// Methods for Movie
func (m *Movie) GetTitle() string    { return m.Title }
func (m *Movie) GetYear() int        { return m.Year }
func (m *Movie) GetRuntime() int     { return int(m.Runtime) }
func (m *Movie) GetGenres() []string { return m.Genres }

// Methods for CreateMovieInput
func (m *CreateMovieInput) GetTitle() string    { return m.Title }
func (m *CreateMovieInput) GetYear() int        { return m.Year }
func (m *CreateMovieInput) GetRuntime() int     { return m.Runtime }
func (m *CreateMovieInput) GetGenres() []string { return m.Genres }

type MovieValidator interface {
	GetTitle() string
	GetYear() int
	GetRuntime() int
	GetGenres() []string
}

func ValidateMovie(v *validator.Validator, m MovieValidator) {
	v.Check(m.GetTitle() != "", "title", "must be provided")
	v.Check(len(m.GetTitle()) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(m.GetYear() != 0, "year", "must be provided")
	v.Check(m.GetYear() >= 1888, "year", "must be greater than 1888")
	v.Check(m.GetYear() <= time.Now().Year(), "year", "must not be in the future")

	v.Check(m.GetRuntime() != 0, "runtime", "must be provided")
	v.Check(m.GetRuntime() > 0, "runtime", "must be a positive integer")

	v.Check(m.GetGenres() != nil, "genres", "must be provided")
	v.Check(len(m.GetGenres()) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(m.GetGenres()) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(m.GetGenres()), "genres", "must not contain duplicate values")
}
