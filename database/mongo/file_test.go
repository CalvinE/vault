package mongo

import (
	"fmt"
	"testing"
	"time"

	"calvinechols.com/vault/file"
)

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
