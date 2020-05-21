package mongo

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

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

// TestMainFile is the test function for this file. its in test main because the testAddFile func needs to run before testGetFile func.
func TestMain(m *testing.M) {
	options := options.Client().ApplyURI(mongoConnectionString)
	client, _ := mongo.Connect(context.TODO(), options)
	fileRepo = NewFileMongoRepo(client)
	accessRepo = NewAccessMongoRepo(client)
	os.Exit(m.Run())
}

func TestAddFile(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestAddFile")
	}
	testFile := &file.File{
		FileID:      fileID,
		CreatedDate: time.Now(),
		Name:        "temp-file-name",
		StorageType: "disk",
		MimeType:    "plain/text",
		WasDeleted:  false,
		OwnerID:     ownerID,
	}
	newFileID, err := fileRepo.AddFile(testFile)
	if err != nil {
		t.Errorf("an error occurred while adding the test file: %v\n", err)
	} else {
		fmt.Printf("new inserted file id: %v\n", newFileID)
	}
}

func TestGetFile(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestGetFile")
	}
	file, err := fileRepo.GetFile(fileID)
	if err != nil {
		t.Errorf("error occurred while getting file from database: %v", err)
	}
	if file.OwnerID != ownerID {
		t.Errorf("the file returned from the does not have the expected ownerID: got = %v expected: %v", file.OwnerID, ownerID)
	} else {
		fmt.Printf("file retreived from database: id = %v, name = %v\n", file.FileID, file.Name)
	}
}
