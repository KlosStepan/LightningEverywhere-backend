package merchant

import (
	"context"
	"errors"
	"lightningeverywhere_backend/internal/models"
)

type Store interface {
	List(ctx context.Context) ([]models.Merchant, error)
	GetByID(ctx context.Context, id string) (*models.Merchant, error)
	Add(ctx context.Context, m *models.Merchant) error
	Update(ctx context.Context, id string, m *models.Merchant) error
	Delete(ctx context.Context, id string) error
}

type memoryStore struct {
	data map[string]models.Merchant
}

func NewMemoryStore() Store {
	m := &memoryStore{data: make(map[string]models.Merchant)}
	for _, merchant := range MockMerchants() {
		m.data[merchant.Properties.ID] = merchant
	}
	return m
}

func (m *memoryStore) List(ctx context.Context) ([]models.Merchant, error) {
	result := make([]models.Merchant, 0, len(m.data))
	for _, merchant := range m.data {
		result = append(result, merchant)
	}
	return result, nil
}

func (m *memoryStore) GetByID(ctx context.Context, id string) (*models.Merchant, error) {
	merchant, ok := m.data[id]
	if !ok {
		return nil, errors.New("merchant not found")
	}
	return &merchant, nil
}

func (m *memoryStore) Add(ctx context.Context, merchant *models.Merchant) error {
	if merchant.Properties.ID == "" {
		return errors.New("merchant ID is required")
	}
	m.data[merchant.Properties.ID] = *merchant
	return nil
}

func (m *memoryStore) Update(ctx context.Context, id string, merchant *models.Merchant) error {
	if _, ok := m.data[id]; !ok {
		return errors.New("merchant not found")
	}
	merchant.Properties.ID = id // enforce ID consistency
	m.data[id] = *merchant
	return nil
}

func (m *memoryStore) Delete(ctx context.Context, id string) error {
	if _, ok := m.data[id]; !ok {
		return errors.New("merchant not found")
	}
	delete(m.data, id)
	return nil
}
