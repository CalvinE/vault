package mongo

import (
	"fmt"
	"testing"
)

func TestAccessRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Mongo Access Tests")
	}

	t.Run("AddAccess", func(t *testing.T) {
		a.Name = accessName
		newAccessID, err := accessRepo.AddAccess(a)
		a.AccessID = newAccessID
		if err != nil {
			t.Errorf("an error occurred while adding the test access: %v\n", err)
		} else {
			fmt.Printf("new inserted accessID: %v\n", newAccessID)
		}
	})

	t.Run("GetAccess", func(t *testing.T) {
		access, err := accessRepo.GetAccess(a.AccessID)
		if err != nil {
			t.Errorf("error occurred while getting access from database: %v", err)
		}
		if access.FileID != f.FileID {
			t.Errorf("the access returned from the does not have the expected FileID: got = %v expected: %v", access.FileID, f.FileID)
		} else {
			fmt.Printf("access retreived from database: id = %v, name = %v\n", access.AccessID, access.Name)
		}
	})

	// TODO Add access log tests
}
