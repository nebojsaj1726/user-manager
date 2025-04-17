package repository

import (
	"context"

	"github.com/nebojsaj1726/user-manager/domain"
	"github.com/nebojsaj1726/user-manager/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c, user)
	return err
}

func (ur *userRepository) Fetch(c context.Context, offset, limit int) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var users []domain.User

	findOptions := options.Find()
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))

	cursor, err := collection.Find(c, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) GetByID(c context.Context, id string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Update(c context.Context, id string, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": user}

	_, err = collection.UpdateOne(c, filter, update)
	return err
}

func (ur *userRepository) Delete(c context.Context, id string) error {
	collection := ur.database.Collection(ur.collection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objID})
	return err
}

func (ur *userRepository) FetchByEmail(c context.Context, email string) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var users []domain.User

	cursor, err := collection.Find(c, bson.M{"email": email})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) Count(c context.Context) (int64, error) {
	collection := ur.database.Collection(ur.collection)
	count, err := collection.CountDocuments(c, bson.M{})
	return count, err
}
