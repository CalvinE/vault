package mongo

import (
	"fmt"
	"testing"

	"calvinechols.com/vault/access"
)

func TestAddAccess(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestAddAccess")
	}
	testAccess := access.NewAccess()
	testAccess.AccessID = accessID
	testAccess.FileID = fileID
	testAccess.Name = accessName
	newAccessID, err := accessRepo.AddAccess(testAccess)
	if err != nil {
		t.Errorf("an error occurred while adding the test access: %v\n", err)
	} else {
		fmt.Printf("new inserted access id: %v\n", newAccessID)
	}
}

func TestGetAccess(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestGetAccess")
	}
	access, err := accessRepo.GetAccess(accessID)
	if err != nil {
		t.Errorf("error occurred while getting access from database: %v", err)
	}
	if access.FileID != fileID {
		t.Errorf("the access returned from the does not have the expected fileID: got = %v expected: %v", access.AccessCount, accessID)
	} else {
		fmt.Printf("access retreived from database: id = %v, name = %v\n", access.AccessID, access.Name)
	}
}
