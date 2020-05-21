package mongo

import (
	"context"

	"calvinechols.com/vault/env"
	"calvinechols.com/vault/file"
	"go.mongodb.org/mongo-driver/bson"
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
		dbName:         env.Get("MONGO_DATABASE", "vault"),
		collectionName: env.Get("MONGO_FILE_COLLECTION", "file"),
	}
}

func (r *fileMongoRepo) AddFile(file *file.File) (string, error) {
	_, err := r.connection.Database(r.dbName).Collection(r.collectionName).InsertOne(context.TODO(), file)
	if err != nil {
		return "", err
	}
	return file.FileID, nil
}

func (r *fileMongoRepo) GetFile(id string) (*file.File, error) {
	// documentID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil, err
	// }
	var file *file.File
	err := r.connection.Database(r.dbName).Collection(r.collectionName).FindOne(context.TODO(), bson.M{"fileId": id}).Decode(&file)
	if err != nil {
		return nil, err
	}
	return file, nil
}
