package mongo

import (
	"fmt"
	"testing"

	"calvinechols.com/vault/file"
)

func TestFileRepo(t *testing.T) {
	f := file.NewFile(fileMimeType, fileName, ownerID)

	t.Run("AddFile", func(t *testing.T) {
		if testing.Short() {
			t.Skip("Skipping AddFile Test")
		}
		newFileID, err := fileRepo.AddFile(f)
		if err != nil {
			t.Errorf("an error occurred while adding the test file: %v\n", err)
		} else {
			fmt.Printf("new inserted file id: %v\n", newFileID)
		}
	})

	// t.Run("GetFile", func(t *testing.T) {
	// 	if testing.Short() {
	// 		t.Skip("Skipping TestGetFile")
	// 	}
	// 	file, err := fileRepo.GetFile(f.FileID)
	// 	if err != nil {
	// 		t.Errorf("error occurred while getting file from database: %v", err)
	// 	}
	// 	if file.FileToken != expectedFileToken {
	// 		t.Errorf("the file returned from the does not have the expected ownerID: got = %v expected: %v", file.FileToken, expectedFileToken)
	// 	} else {
	// 		fmt.Printf("file retreived from database: id = %v, name = %v\n", file.FileID, file.Name)
	// 	}
	// })

	t.Run("GetFileByFileToken", func(t *testing.T) {
		if testing.Short() {
			t.Skip("Skipping GetFileByFileToken Test")
		}
		file, err := fileRepo.GetFileByFileToken(f.FileToken)
		if err != nil {
			t.Errorf("error occurred while getting file from database: %v", err)
		}
		if file.OwnerID != ownerID {
			t.Errorf("the file returned from the does not have the expected ownerID: got = %v expected: %v", file.OwnerID, ownerID)
		} else {
			fmt.Printf("file retreived from database: id = %v, name = %v\n", file.FileID, file.Name)
		}
	})
}
