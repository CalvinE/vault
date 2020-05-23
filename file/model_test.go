package file

import (
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	mimeType, fileName, ownerID, storageType := "test1", "test2", "test3", "test4"
	f := NewFile(mimeType, fileName, ownerID, storageType)
	unsetTime := time.Time{}
	if f.MimeType != mimeType {
		t.Errorf("MimeType is wrong: got: %v - expected: %v", f.MimeType, mimeType)
	}
	if f.Name != fileName {
		t.Errorf("Name is wrong: got: %v - expected: %v", f.Name, fileName)
	}
	if f.OwnerID != ownerID {
		t.Errorf("OwnerID is wrong: got: %v - expected: %v", f.OwnerID, ownerID)
	}
	if f.StorageType != storageType {
		t.Errorf("StorageType is wrong: got: %v - expected: %v", f.StorageType, storageType)
	}
	if f.CreatedDate == unsetTime {
		t.Errorf("CreatedDate should be initialized to the time NewFile is called.\n")
	}
}

func TestValidate(t *testing.T) {
	var file *File
	err := file.Validate()
	if err == nil {
		t.Error("a nil file should not be valid.\n")
	}
	// TODO: finish this test...
}

func TestIsDeleted(t *testing.T) {
	file := NewFile("", "", "", "")
	isDeleted := file.IsDeleted()
	if isDeleted == true {
		t.Errorf("file without DeletedDate set should return false got: %v\n", isDeleted)
	}
	file.DeletedDate = time.Now()
	isDeleted = file.IsDeleted()
	if isDeleted == false {
		t.Errorf("file with DeletedDate set should return true got: %v\n", isDeleted)
	}
}

func TestIsExpired(t *testing.T) {
	file := NewFile("", "", "", "")
	isExpired := file.IsExpired()
	if isExpired == true {
		t.Errorf("ExpirationDate being unset should result in IsExpired returning false: ExpirationDate: %v, isExpired: %v", file.ExpirationDate, isExpired)
	}
	file.ExpirationDate = time.Now().Add(-10000)
	isExpired = file.IsExpired()
	if isExpired == false {
		t.Errorf("ExpirationDate in the past should result in IsExpired returning true: ExpirationDate: %v, isExpired: %v\n", file.ExpirationDate, isExpired)
	}
	file.ExpirationDate = time.Now().Add(100)
	isExpired = file.IsExpired()
	if isExpired == true {
		t.Errorf("ExpirationDate in the future should result in IsExpired returning false: ExpirationDate: %v, isExpired: %v", file.ExpirationDate, isExpired)
	}
}
