package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	param := httprouter.ParamsFromContext(r.Context()).ByName("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if id <= 0 || err != nil{
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
