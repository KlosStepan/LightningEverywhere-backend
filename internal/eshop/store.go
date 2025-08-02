package eshop

import (
	"context"
	//"errors"
	//"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

// MongoStore implements Store using MongoDB
type mongoStore struct {
	coll *mongo.Collection
}

// NewMongoStore creates a new Mongo store instance
func NewMongoStore(db *mongo.Database) Store {
	return &mongoStore{
		coll: db.Collection("eshops"),
	}
}

func (m *mongoStore) ReadAll(ctx context.Context) ([]Eshop, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var eshops []Eshop
	if err := cursor.All(ctx, &eshops); err != nil {
		return nil, err
	}
	return eshops, nil
}

func (m *mongoStore) Read(ctx context.Context, id string) (*Eshop, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var e Eshop
	err := m.coll.FindOne(ctx, bson.M{"id": id}).Decode(&e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (m *mongoStore) Create(ctx context.Context, e *Eshop) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.coll.InsertOne(ctx, e)
	return err
}

func (m *mongoStore) Update(ctx context.Context, id string, e *Eshop) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.coll.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": e})
	return err
}

func (m *mongoStore) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.coll.DeleteOne(ctx, bson.M{"id": id})
	return err
}
