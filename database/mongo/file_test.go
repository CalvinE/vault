package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"calvinechols.com/vault/env"
	"calvinechols.com/vault/file"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestMainFile is the test function for this file. its in test main because the testAddFile func needs to run before testGetFile func.
func TestMainFile(t *testing.T) {
	fileID, ownerID, mongoConnectionString := uuid.NewV4().String(), uuid.NewV4().String(), env.Get("MONGODB_DATA_CONNECTION", "mongodb://root:password@localhost:27017")
	options := options.Client().ApplyURI(mongoConnectionString)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		t.Errorf("an error occurred while connecting to database: %v", err)
	}
	fileMongoRepo := NewFileMongoRepo(client)
	testAddFile(t, fileMongoRepo, fileID, ownerID)
	testGetFile(t, fileMongoRepo, fileID, ownerID)
}

func testAddFile(t *testing.T, fileMongoRepo file.Repo, fileID, fileOwnerID string) {

	testFile := &file.File{
		FileID:      fileID,
		CreatedDate: time.Now(),
		Name:        "temp-file-name",
		StorageType: "disk",
		MimeType:    "plain/text",
		WasDeleted:  false,
		OwnerID:     fileOwnerID,
	}
	newFileID, err := fileMongoRepo.AddFile(testFile)
	if err != nil {
		t.Errorf("an error occurred while adding the test file: %v", err)
	} else {
		fmt.Printf("new inserted file database id: %v\n", newFileID)
	}

}

func testGetFile(t *testing.T, fileMongoRepo file.Repo, fileID, fileOwnerID string) {
	file, err := fileMongoRepo.GetFile(fileID)
	if err != nil {
		t.Errorf("error occurred while getting file from database: %v", err)
	}
	if file.OwnerID != fileOwnerID {
		t.Errorf("the file returned from the does not have the expected ownerID: got = %v expected: %v", file.OwnerID, fileOwnerID)
	} else {
		fmt.Printf("file retreived from database: id = %v, name = %v\n", file.FileID, file.Name)
	}
}
