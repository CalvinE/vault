package access

import (
	"testing"
	"time"
)

func TestNewAccess(t *testing.T) {
	a := NewAccess()
	unsetTime := time.Time{}
	if a.AccessID == "" {
		t.Error("newly created access id is empty.\n")
	}
	if a.AccessCount != 0 {
		t.Error("newly created access AccessCount should initialize to 0.\n")
	}
	if a.CreatedDate == unsetTime {
		t.Error("newly created access CreatedDate to be initialized to the time the access object is instantiated.\n")
	}
}

func TestVaildate(t *testing.T) {
	var access *Access
	err := access.Validate()
	if err == nil {
		t.Error("a nil access should not be valid.\n")
	}
	// TODO: finish this test...
}

func TestIsDisabled(t *testing.T) {
	access := NewAccess()
	isDeleted := access.IsDisabled()
	if isDeleted == true {
		t.Errorf("file without DisabledDate set should return false got: %v\n", isDeleted)
	}
	access.DisabledDate = time.Now()
	isDeleted = access.IsDisabled()
	if isDeleted == false {
		t.Errorf("file with DisabledDate set should return true got: %v\n", isDeleted)
	}

}

func TestIsExpired(t *testing.T) {
	access := NewAccess()
	isExpired := access.IsExpired()
	if isExpired == true {
		t.Errorf("ExpirationDate being unset should result in IsExpired returning false: ExpirationDate: %v, isExpired: %v", access.ExpirationDate, isExpired)
	}
	access.ExpirationDate = time.Now().Add(-10000)
	isExpired = access.IsExpired()
	if isExpired == false {
		t.Errorf("ExpirationDate in the past should result in IsExpired returning true: ExpirationDate: %v, isExpired: %v\n", access.ExpirationDate, isExpired)
	}
	access.ExpirationDate = time.Now().Add(100)
	isExpired = access.IsExpired()
	if isExpired == true {
		t.Errorf("ExpirationDate in the future should result in IsExpired returning false: ExpirationDate: %v, isExpired: %v", access.ExpirationDate, isExpired)
	}
}
