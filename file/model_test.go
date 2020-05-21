package file

import (
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	f := NewFile()
	unsetTime := time.Time{}
	if f.FileID == "" {
		t.Error("newly created file id is empty.\n")
	}
	if f.CreatedDate == unsetTime {
		t.Errorf("CreatedDate should be initialized to the time NewFile is called\n")
	}
}

func TestValidate(t *testing.T) {
	var file *File
	isValid, errorMessages := file.Validate()
	if isValid == true {
		t.Error("a nil file is not valid\n")
	}
	numErrorMessages := len(errorMessages)
	if numErrorMessages != 1 {
		t.Error("there should be one error message for a nil file\n")
	}
}

func TestIsDeleted(t *testing.T) {
	file := NewFile()
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
	file := NewFile()
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
		t.Errorf("ExpirationDate in the future should resulkt in IsExpired returning false: ExpirationDate: %v, isExpired: %v", file.ExpirationDate, isExpired)
	}
}
