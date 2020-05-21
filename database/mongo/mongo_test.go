package mongo

import (
	"context"
	"os"
	"testing"

	"calvinechols.com/vault/access"
	"calvinechols.com/vault/env"
	"calvinechols.com/vault/file"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoConnectionString = env.Get("MONGODB_DATA_CONNECTION", "mongodb://root:password@localhost:27017")

var fileID = uuid.NewV4().String()
var ownerID = uuid.NewV4().String()
var fileRepo file.Repo

var accessID = uuid.NewV4().String()
var accessName = "test-access"
var accessRepo access.Repo

// TestMainFile is the test function for this file. its in test main because the testAddFile func needs to run before testGetFile func.
func TestMain(m *testing.M) {
	options := options.Client().ApplyURI(mongoConnectionString)
	client, _ := mongo.Connect(context.TODO(), options)
	fileRepo = NewFileMongoRepo(client)
	accessRepo = NewAccessMongoRepo(client)
	os.Exit(m.Run())
}
