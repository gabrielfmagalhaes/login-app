package repository

import (
	"context"
	"login-app/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Insert(user *domain.User) (string, error)
	FindByEmail(email string) (domain.User, error)
}

type mongoRepository struct {
	collectionName string
	collection     *mongo.Collection
}

func NewRepository(c *mongo.Collection) Repository {
	return &mongoRepository{collection: c}
}

func (r *mongoRepository) Insert(user *domain.User) (string, error) {
	result, err := r.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).String(), err
}

func (r *mongoRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	filter := bson.D{{"email", email}}

	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)

	return user, err
}
