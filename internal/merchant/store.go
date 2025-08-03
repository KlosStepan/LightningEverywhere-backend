package merchant

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	ReadAll(ctx context.Context) ([]Merchant, error)
	Create(ctx context.Context, m *Merchant) error
	Read(ctx context.Context, id string) (*Merchant, error)
	Update(ctx context.Context, id string, m *Merchant) error
	Delete(ctx context.Context, id string) error
}

// mongoStore implements Store using MongoDB
type mongoStore struct {
	coll *mongo.Collection
}

// NewMongoStore creates a new MongoDB-backed store for merchants
func NewMongoStore(db *mongo.Database) Store {
	return &mongoStore{
		coll: db.Collection("merchants"),
	}
}

func (m *mongoStore) ReadAll(ctx context.Context) ([]Merchant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var merchants []Merchant
	if err := cursor.All(ctx, &merchants); err != nil {
		return nil, err
	}
	return merchants, nil
}

func (m *mongoStore) Read(ctx context.Context, id string) (*Merchant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var merchant Merchant
	err := m.coll.FindOne(ctx, bson.M{"properties.id": id}).Decode(&merchant)
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (m *mongoStore) Create(ctx context.Context, merchant *Merchant) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.coll.InsertOne(ctx, merchant)
	return err
}

func (m *mongoStore) Update(ctx context.Context, id string, merchant *Merchant) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.coll.UpdateOne(ctx, bson.M{"properties.id": id}, bson.M{"$set": merchant})
	return err
}

func (m *mongoStore) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.coll.DeleteOne(ctx, bson.M{"properties.id": id})
	return err
}
