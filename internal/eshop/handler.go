package eshop

import (
	//"encoding/json"
	"net/http"
)

func MakeHandler( /*store Store*/ ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			//GetAll(w, r, store)
		case http.MethodPost:
			//Create(w, r, store)
		case http.MethodPut:
			//Update(w, r, store)
		case http.MethodDelete:
			//Delete(w, r, store)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
