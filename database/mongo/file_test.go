package mongo

import (
	"fmt"
	"testing"
)

func TestFileRepo(t *testing.T) {
	// f := file.NewFile(fileMimeType, fileName, ownerID)

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

	t.Run("GetFile", func(t *testing.T) {
		if testing.Short() {
			t.Skip("Skipping TestGetFile")
		}
		file, err := fileRepo.GetFile(f.FileID)
		if err != nil {
			t.Errorf("error occurred while getting file from database: %v", err)
		}
		if file.Name != fileName {
			t.Errorf("the file returned from the does not have the expected ownerID: got = %v expected: %v", file.Name, fileName)
		} else {
			fmt.Printf("file retreived from database: id = %v, name = %v\n", file.FileID, file.Name)
		}
	})
}
