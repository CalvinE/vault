package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"calvinechols.com/vault/file"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAddFile(t *testing.T) {
	options := options.Client().ApplyURI("mongodb://root:password@localhost:27017")
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		t.Errorf("An error occurred while connecting to database: %v", err)
	}
	fileMongoRepo := NewFileMongoRepo(client)
	fileID, ownerID := uuid.NewV4().String(), uuid.NewV4().String()
	testFile := &file.DBFile{
		File: file.File{
			FileID:      fileID,
			CreatedDate: time.Now(),
			FileName:    "temp-file-name",
			StorageType: "disk",
			MimeType:    "plain/text",
			WasDeleted:  false,
			OwnerID:     ownerID,
		},
	}
	newFileID, err := fileMongoRepo.AddFile(testFile)
	if err != nil {
		t.Errorf("an error occurred while adding the test file: %v", err)
	} else {
		fmt.Printf("New inserted file database id: %v", newFileID)
	}
	fmt.Printf("new file added to database with ObjectId: %v", newFileID)

}

func TestGetFile(t *testing.T) {

}
