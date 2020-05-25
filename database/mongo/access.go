package mongo

import (
	"context"

	"calvinechols.com/vault/access"
	"calvinechols.com/vault/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type accessMongoRepo struct {
	connection           *mongo.Client
	dbName               string
	accessCollectionName string
	logCollectionName    string
}

// NewAccessMongoRepo returns a new instance of the accessMongoRepo
func NewAccessMongoRepo(connection *mongo.Client) access.Repo {
	return &accessMongoRepo{
		connection:           connection,
		dbName:               env.Get("MONGO_DATABASE", "vault"),
		accessCollectionName: env.Get("MONGO_ACCESS_COLLECTION", "access"),
		logCollectionName:    env.Get("MONGO_LOG_COLLECTION", "log"),
	}
}

func (r *accessMongoRepo) AddAccess(access *access.Access) (primitive.ObjectID, error) {
	insertResult, err := r.connection.Database(r.dbName).Collection(r.accessCollectionName).InsertOne(context.TODO(), access)
	if err != nil {
		return primitive.NilObjectID, err
	}
	access.AccessID = insertResult.InsertedID.(primitive.ObjectID)
	return access.AccessID, nil
}

func (r *accessMongoRepo) GetAccess(accessID primitive.ObjectID) (*access.Access, error) {
	var access *access.Access
	err := r.connection.Database(r.dbName).Collection(r.accessCollectionName).FindOne(context.TODO(), bson.M{"_id": accessID}).Decode(&access)
	if err != nil {
		return nil, err
	}
	return access, nil
}

func (r *accessMongoRepo) AddLog(log *access.Log) (primitive.ObjectID, error) {
	insertResult, err := r.connection.Database(r.dbName).Collection(r.logCollectionName).InsertOne(context.TODO(), log)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func (r *accessMongoRepo) GetLogsByAccessID(accessID primitive.ObjectID) ([]access.Log, error) {
	var logs []access.Log
	findResult, err := r.connection.Database(r.dbName).Collection(r.logCollectionName).Find(context.TODO(), bson.M{"accessId": accessID})
	if err != nil {
		return nil, err
	}
	err = findResult.All(context.TODO(), &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *accessMongoRepo) GetLogsByFileID(fileID primitive.ObjectID) ([]access.Log, error) {
	var logs []access.Log
	findResult, err := r.connection.Database(r.dbName).Collection(r.logCollectionName).Aggregate(context.TODO(), bson.M{"fileId": fileID})
	if err != nil {
		return nil, err
	}
	err = findResult.All(context.TODO(), &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
