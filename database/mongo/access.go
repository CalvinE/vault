package mongo

import (
	"context"

	"calvinechols.com/vault/access"
	"calvinechols.com/vault/env"
	"go.mongodb.org/mongo-driver/bson"
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
		dbName:         env.Get("MONGO_DATABASE", "vault"),
		collectionName: env.Get("MONGO_ACCESS_COLLECTION", "access"),
	}
}

func (r *accessMongoRepo) AddAccess(access *access.Access) (string, error) {
	_, err := r.connection.Database(r.dbName).Collection(r.collectionName).InsertOne(context.TODO(), access)
	if err != nil {
		return "", err
	}
	return access.AccessID, nil
}

func (r *accessMongoRepo) GetAccess(id string) (*access.Access, error) {
	var access *access.Access
	err := r.connection.Database(r.dbName).Collection(r.collectionName).FindOne(context.TODO(), bson.M{"accessId": id}).Decode(&access)
	if err != nil {
		return nil, err
	}
	return access, nil
}
