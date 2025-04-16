package merchant

import "lightningeverywhere_backend/internal/models"

type Store interface {
	List() []models.Merchant
	GetByID(id string) *models.Merchant
	Add(merchant models.Merchant) error
	Update(id string, merchant models.Merchant) error
	Delete(id string) error
}

type mockStore struct {
	data map[string]models.Merchant
}

func NewMockStore() Store {
	store := &mockStore{
		data: make(map[string]models.Merchant),
	}
	for _, m := range dummyMerchants {
		store.data[m.Properties.ID] = m
	}
	return store
}

func (s *mockStore) List() []models.Merchant {
	merchants := []models.Merchant{}
	for _, m := range s.data {
		merchants = append(merchants, m)
	}
	return merchants
}

func (s *mockStore) GetByID(id string) *models.Merchant {
	merchant, exists := s.data[id]
	if !exists {
		return nil
	}
	return &merchant
}

func (s *mockStore) Add(m models.Merchant) error {
	s.data[m.Properties.ID] = m
	return nil
}

func (s *mockStore) Update(id string, m models.Merchant) error {
	if _, ok := s.data[id]; !ok {
		return nil
	}
	s.data[id] = m
	return nil
}

func (s *mockStore) Delete(id string) error {
	delete(s.data, id)
	return nil
}
