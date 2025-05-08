package eshop

import (
	"context"
	"errors"
	"fmt"
)

type Store interface {
	ReadAll(ctx context.Context) ([]Eshop, error)
	Create(ctx context.Context, e *Eshop) error
	Read(ctx context.Context, id string) (*Eshop, error) // ✅ Added method to interface
	Update(ctx context.Context, id string, e *Eshop) error
	Delete(ctx context.Context, id string) error
}

type memoryStore struct {
	data map[string]Eshop
}

// NewMemoryStore creates and returns a new memory-based store, initializing it with mock eshops
func NewMemoryStore() Store {
	m := &memoryStore{data: make(map[string]Eshop)}
	for _, e := range MockEshops() {
		m.data[e.ID] = e
	}
	return m
}

func (m *memoryStore) ReadAll(ctx context.Context) ([]Eshop, error) {
	result := make([]Eshop, 0, len(m.data))
	for _, e := range m.data {
		result = append(result, e)
	}
	return result, nil
}

func (m *memoryStore) Read(ctx context.Context, id string) (*Eshop, error) {
	e, ok := m.data[id]
	if !ok {
		return nil, errors.New("eshop not found")
	}
	return &e, nil
}

func (m *memoryStore) Create(ctx context.Context, e *Eshop) error {
	if e.ID == "" {
		//e.ID = uuid.New().String()
	}
	m.data[e.ID] = *e
	return nil
}

func (m *memoryStore) Update(ctx context.Context, id string, e *Eshop) error {
	if _, ok := m.data[id]; !ok {
		return fmt.Errorf("eshop not found")
	}
	e.ID = id // ensure the ID stays consistent
	m.data[id] = *e
	return nil
}

func (m *memoryStore) Delete(ctx context.Context, id string) error {
	if _, ok := m.data[id]; !ok {
		return fmt.Errorf("eshop not found")
	}
	delete(m.data, id)
	return nil
}
