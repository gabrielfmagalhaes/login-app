package mongodb

import (
	"context"
	"login-app/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(c *mongo.Collection) *Repository {
	return &Repository{collection: c}
}

func (r *Repository) Insert(user *domain.User) (string, error) {
	result, err := r.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).String(), nil
}
