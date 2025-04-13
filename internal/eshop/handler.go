package eshop

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakeHandler(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAll(w, r, store)
		case http.MethodPost:
			Create(w, r, store)
		case http.MethodPut:
			Update(w, r, store)
		case http.MethodDelete:
			Delete(w, r, store)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func GetAll(w http.ResponseWriter, r *http.Request, store Store) {
	w.Header().Set("Content-Type", "application/json")
	eshops, err := store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to get eshops", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(eshops)
}

func Create(w http.ResponseWriter, r *http.Request, store Store) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var e Eshop
	if err := json.Unmarshal(body, &e); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := store.Create(r.Context(), &e); err != nil {
		http.Error(w, fmt.Sprintf("failed to create eshop: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func Update(w http.ResponseWriter, r *http.Request, store Store) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var e Eshop
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := store.Update(r.Context(), id, &e); err != nil {
		http.Error(w, fmt.Sprintf("failed to update eshop: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func Delete(w http.ResponseWriter, r *http.Request, store Store) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := store.Delete(r.Context(), id); err != nil {
		http.Error(w, fmt.Sprintf("failed to delete eshop: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
