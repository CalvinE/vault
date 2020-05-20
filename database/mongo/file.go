package database

import (
	"context"

	"calvinechols.com/vault/file"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type fileMongoRepo struct {
	connection     *mongo.Client
	dbName         string
	collectionName string
}

// NewFileMongoRepo returns a new instance of the fileMongoRepo struct
func NewFileMongoRepo(connection *mongo.Client) file.Repo {
	return &fileMongoRepo{
		connection:     connection,
		dbName:         "vault",
		collectionName: "file",
	}
}

func (r *fileMongoRepo) AddFile(file *file.DBFile) (string, error) {
	result, err := r.connection.Database(r.dbName).Collection(r.collectionName).InsertOne(context.TODO(), file)
	if err != nil {
		return "", err
	}
	idString := result.InsertedID.(primitive.ObjectID).String()
	return idString, nil
}

func (r *fileMongoRepo) GetFile(id string) (*file.DBFile, error) {
	documentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var file *file.DBFile
	err = r.connection.Database(r.dbName).Collection(r.collectionName).FindOne(context.TODO(), bson.M{"_id": documentID}).Decode(&file)
	if err != nil {
		return nil, err
	}
	return file, nil
}
