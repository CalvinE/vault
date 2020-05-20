package database

import (
	"context"

	"calvinechols.com/vault/access"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type accessMongoRepo struct {
	connection     *mongo.Client
	dbName         string
	collectionName string
}

// NewAccessMongoRepo returns a new instance of the accessMongoRepo
func NewAccessMongoRepo(connection *mongo.Client) access.Repo {
	return &accessMongoRepo{
		connection:     connection,
		dbName:         "vault",
		collectionName: "access",
	}
}

func (r *accessMongoRepo) AddAccess(access *access.DBAccess) (string, error) {
	result, err := r.connection.Database(r.dbName).Collection(r.collectionName).InsertOne(context.TODO(), access)
	if err != nil {
		return "", err
	}
	idString := result.InsertedID.(primitive.ObjectID).String()
	return idString, nil
}

func (r *accessMongoRepo) GetAccess(id string) (*access.DBAccess, error) {
	documentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var access *access.DBAccess
	err = r.connection.Database(r.dbName).Collection(r.collectionName).FindOne(context.TODO(), bson.M{"_id": documentID}).Decode(&access)
	if err != nil {
		return nil, err
	}
	return access, nil
}
