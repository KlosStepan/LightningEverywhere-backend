package merchant

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
			id := r.URL.Query().Get("id")
			if id != "" {
				Read(w, r, store) // Single
			} else {
				ReadAll(w, r, store) // All
			}
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

func ReadAll(w http.ResponseWriter, r *http.Request, store Store) {
	w.Header().Set("Content-Type", "application/json")
	merchants, err := store.ReadAll(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch merchants", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(merchants)
}

func Create(w http.ResponseWriter, r *http.Request, store Store) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var m Merchant
	if err := json.Unmarshal(body, &m); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := store.Create(r.Context(), &m); err != nil {
		http.Error(w, fmt.Sprintf("failed to create merchant: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func Read(w http.ResponseWriter, r *http.Request, store Store) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	merchant, err := store.Read(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch merchant: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(merchant)
}

func Update(w http.ResponseWriter, r *http.Request, store Store) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var m Merchant
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := store.Update(r.Context(), id, &m); err != nil {
		http.Error(w, fmt.Sprintf("failed to update merchant: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func Delete(w http.ResponseWriter, r *http.Request, store Store) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := store.Delete(r.Context(), id); err != nil {
		http.Error(w, fmt.Sprintf("failed to delete merchant: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
