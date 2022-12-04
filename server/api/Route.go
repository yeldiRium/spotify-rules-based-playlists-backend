package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route func(http.ResponseWriter, *http.Request, httprouter.Params) *Error

func (fn Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	if e := fn(w, r, params); e != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(e.Code)

		marshalledError, err := json.Marshal(e)
		if err != nil {
			panic(err)
		}

		w.Write(marshalledError)
	}
}
