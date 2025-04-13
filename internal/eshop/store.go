package eshop

import (
	"context"
	"fmt"
)

type memoryStore struct {
	data map[string]Eshop
}

type Store interface {
	GetAll(ctx context.Context) ([]Eshop, error)
	Create(ctx context.Context, e *Eshop) error
	Update(ctx context.Context, id string, e *Eshop) error
	Delete(ctx context.Context, id string) error
}

func NewMemoryStore() Store {
	m := &memoryStore{data: make(map[string]Eshop)}
	for _, e := range MockEshops() {
		m.data[e.ID] = e
	}
	return m
}

func (m *memoryStore) GetAll(ctx context.Context) ([]Eshop, error) {
	result := make([]Eshop, 0, len(m.data))
	for _, e := range m.data {
		result = append(result, e)
	}
	return result, nil
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
