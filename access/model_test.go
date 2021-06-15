package access

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewAccess(t *testing.T) {
	var creatorID = uuid.NewV4().String()
	a := NewAccess(primitive.NewObjectID(), creatorID)
	if a.AccessCount != 0 {
		t.Error("newly created access AccessCount should initialize to 0.\n")
	}
	if a.CreatedDate == nil {
		t.Error("newly created access CreatedDate to be initialized to the time the access object is instantiated.\n")
	}
	if a.CreatorID != creatorID {
		t.Errorf("CreatorID is wrong value: expected: %v - got: %v\n", creatorID, a.CreatorID)
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
	access := NewAccess(primitive.NewObjectID(), uuid.NewV4().String())
	isDeleted := access.IsDisabled()
	if isDeleted == true {
		t.Errorf("file without DisabledDate set should return false got: %v\n", isDeleted)
	}
	now := time.Now()
	access.DisabledDate = &now
	isDeleted = access.IsDisabled()
	if isDeleted == false {
		t.Errorf("file with DisabledDate set should return true got: %v\n", isDeleted)
	}

}

func TestIsExpired(t *testing.T) {
	access := NewAccess(primitive.NewObjectID(), uuid.NewV4().String())
	isExpired := access.IsExpired()
	if isExpired == true {
		t.Errorf("ExpirationDate being unset should result in IsExpired returning false: ExpirationDate: %v, isExpired: %v", access.ExpirationDate, isExpired)
	}
	pastExp := time.Now().Add(time.Minute * -1)
	access.ExpirationDate = &pastExp
	isExpired = access.IsExpired()
	if isExpired == false {
		t.Errorf("ExpirationDate in the past should result in IsExpired returning true: ExpirationDate: %v, isExpired: %v\n", access.ExpirationDate, isExpired)
	}
	futureExp := time.Now().Add(time.Minute * 1)
	access.ExpirationDate = &futureExp
	isExpired = access.IsExpired()
	if isExpired == true {
		t.Errorf("ExpirationDate in the future should result in IsExpired returning false: ExpirationDate: %v, isExpired: %v", access.ExpirationDate, isExpired)
	}
}
