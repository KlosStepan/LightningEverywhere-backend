package merchant

import (
	"context"
	"errors"
)

type Store interface {
	ReadAll(ctx context.Context) ([]Merchant, error)
	Create(ctx context.Context, m *Merchant) error
	Read(ctx context.Context, id string) (*Merchant, error)
	Update(ctx context.Context, id string, m *Merchant) error
	Delete(ctx context.Context, id string) error
}

type memoryStore struct {
	data map[string]Merchant
}

// NewMemoryStore creates and returns a new memory-based store, initializing it with mock merchants
func NewMemoryStore() Store {
	m := &memoryStore{data: make(map[string]Merchant)}
	for _, merchant := range MockMerchants() {
		m.data[merchant.Properties.ID] = merchant
	}
	return m
}

func (m *memoryStore) ReadAll(ctx context.Context) ([]Merchant, error) {
	result := make([]Merchant, 0, len(m.data))
	for _, merchant := range m.data {
		result = append(result, merchant)
	}
	return result, nil
}

// Read retrieves a merchant by its ID from the store
func (m *memoryStore) Read(ctx context.Context, id string) (*Merchant, error) {
	merchant, ok := m.data[id]
	if !ok {
		return nil, errors.New("merchant not found")
	}
	return &merchant, nil
}

// Add adds a new merchant to the store
func (m *memoryStore) Create(ctx context.Context, merchant *Merchant) error {
	if merchant.Properties.ID == "" {
		return errors.New("merchant ID is required")
	}
	m.data[merchant.Properties.ID] = *merchant
	return nil
}

// Update updates an existing merchant in the store by its ID
func (m *memoryStore) Update(ctx context.Context, id string, merchant *Merchant) error {
	if _, ok := m.data[id]; !ok {
		return errors.New("merchant not found")
	}
	merchant.Properties.ID = id // enforce ID consistency
	m.data[id] = *merchant
	return nil
}

// Delete removes a merchant from the store by its ID
func (m *memoryStore) Delete(ctx context.Context, id string) error {
	if _, ok := m.data[id]; !ok {
		return errors.New("merchant not found")
	}
	delete(m.data, id)
	return nil
}
