package mongo

import (
	"context"
	"testing"

	"calvinechols.com/vault/access"
	"calvinechols.com/vault/env"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestMainAccess is the test function for this file. its in test main because the testAddAccess func needs to run before testGetAccess func.
func TestMainAccess(t *testing.T) {
	accessID, fileID, mongoConnectionString := uuid.NewV4().String(), uuid.NewV4().String(), env.Get("MONGODB_DATA_CONNECTION", "mongodb://root:password@localhost:27017")
	options := options.Client().ApplyURI(mongoConnectionString)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		t.Errorf("an error occurred while connecting to database: %v", err)
	}
	accessMongoRepo := NewAccessMongoRepo(client)
	testAddAccess(t, accessMongoRepo, accessID, fileID)
	testGetAccess(t, accessMongoRepo, accessID, fileID)
}

func testAddAccess(t *testing.T, accessMongoRepo access.Repo, accessID, fileID string) {

}

func testGetAccess(t *testing.T, accessMongoRepo access.Repo, accessID, fileID string) {

}
