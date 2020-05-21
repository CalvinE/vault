package file

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// File is a type representing a file in the vault.
type File struct {
	FileID         string    `json:"fileId" bson:"fileId"`
	CreatedDate    time.Time `json:"createdDate" bson:"createdDate"`
	StorageType    string    `json:"storageType" bson:"storageType"`
	Name           string    `json:"name" bson:"name"`
	MimeType       string    `json:"mimeType" bson:"mimeType"`
	DeletedDate    time.Time `json:"dateDeleted,omitempty" bson:"dateDeleted,omitempty"`
	ExpirationDate time.Time `json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	OwnerID        string    `json:"ownerId" bson:"ownerId"`
}

// DBFile represents a file from the database
type DBFile struct {
	File `bson:",inline"`
	DbID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

// NewFile returns a new File object
func NewFile() *File {
	return &File{
		FileID:      uuid.NewV4().String(),
		CreatedDate: time.Now(),
	}
}

// Validate is a functi on that validates a File struct to make sure it has valid data.
func (f *File) Validate() (bool, map[string]string) {
	isValid := false
	errorMessages := make(map[string]string)

	if f == nil {
		isValid = false
		errorMessages["General"] = "the file cannot be nil."
		return isValid, errorMessages
	}

	if f.MimeType == "" {
		isValid = false
		errorMessages["MimeType"] = fmt.Sprint("MimeType cannot be empty")
	}

	if f.StorageType == "" {
		isValid = false
		errorMessages["StorageType"] = fmt.Sprint("StorageType cannot be empty")
	}

	if f.OwnerID == "" {
		isValid = false
		errorMessages["OwnerID"] = fmt.Sprint("OwnerID cannot be empty")
	}

	if f.Name == "" {
		isValid = false
		errorMessages["Name"] = fmt.Sprint("Name cannot be empty")
	}

	return isValid, errorMessages
}

// IsDeleted returns true if the file's deletedDate is set.
func (f *File) IsDeleted() bool {
	unsetTime := time.Time{}
	wasDeleted := f.DeletedDate != unsetTime
	return wasDeleted
}

// IsExpired returns true if the file has an expiration data and it is after the current time.
func (f *File) IsExpired() bool {
	unsetTime := time.Time{}
	hasExpirationDate := f.ExpirationDate != unsetTime
	if hasExpirationDate == true {
		now := time.Now()
		isExpired := now.After(f.ExpirationDate)
		return isExpired
	}
	return hasExpirationDate

}
