package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	param := httprouter.ParamsFromContext(r.Context()).ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if id <= 0 || err != nil {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	jsn, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Newline to make viewing easier in terminal applications.
	jsn = append(jsn, '\n')

	// It's OK if the provided header map is nil. Go doesn't throw an error
	// if you try to range over (or generally, read from) a nil map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsn)
	return nil
}
