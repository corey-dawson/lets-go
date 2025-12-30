package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.letsgo.com/internals/models"
	"greenlight.letsgo.com/internals/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input models.CreateMovieInput

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Initialize a new Validator.
	v := validator.New()
	// Call the ValidateMovie() function, and if any checks fail, return a response
	if models.ValidateMovie(v, &input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := models.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// fmt.Fprintf(w, "show the details of movie %d\n", id)

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}
