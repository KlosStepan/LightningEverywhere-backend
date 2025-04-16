package merchant

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	merchants, err := h.store.GetAll()
	if err != nil {
		http.Error(w, "failed to fetch merchants", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(merchants)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	merchant, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, "merchant not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(merchant)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var m Merchant
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	if err := h.store.Create(m); err != nil {
		http.Error(w, "could not create merchant", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var m Merchant
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	if err := h.store.Update(id, m); err != nil {
		http.Error(w, "could not update merchant", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.store.Delete(id); err != nil {
		http.Error(w, "could not delete merchant", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
