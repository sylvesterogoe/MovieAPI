package main

import (
	"fmt"
	"net/http"
	"time"

	"starlet.sylvester.net/internal/data"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Akacalito",
		Year:      2023,
		Runtime:   120,
		Genres:    []string{"drama", "comedy", "adventure"},
		Version:   1,
	}

	app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "movie created")
}
