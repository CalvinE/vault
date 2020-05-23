package mongo

import (
	"fmt"
	"testing"
)

func TestAccessRepo(t *testing.T) {
	t.Run("AddAccess", func(t *testing.T) {
		if testing.Short() {
			t.Skip("Skipping AddAccess Test")
		}
		a.Name = accessName
		newAccessToken, err := accessRepo.AddAccess(a)
		if err != nil {
			t.Errorf("an error occurred while adding the test access: %v\n", err)
		} else {
			fmt.Printf("new inserted access token: %v\n", newAccessToken)
		}
	})

	t.Run("GetAccessByAccessToken", func(t *testing.T) {
		if testing.Short() {
			t.Skip("Skipping GetAccess Test")
		}
		access, err := accessRepo.GetAccessByAccessToken(a.AccessToken)
		if err != nil {
			t.Errorf("error occurred while getting access from database: %v", err)
		}
		if access.FileToken != f.FileToken {
			t.Errorf("the access returned from the does not have the expected fileToken: got = %v expected: %v", access.FileToken, f.FileToken)
		} else {
			fmt.Printf("access retreived from database: id = %v, name = %v\n", access.AccessID, access.Name)
		}
	})

	// TODO Add access log tests
}
