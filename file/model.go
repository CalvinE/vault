package file

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// File is a type representing a file in the vault.
type File struct {
	FileID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate    time.Time          `json:"createdDate" bson:"createdDate"`
	StorageType    string             `json:"storageType" bson:"storageType"`
	Name           string             `json:"name" bson:"name"`
	MimeType       string             `json:"mimeType" bson:"mimeType"`
	DeletedDate    time.Time          `json:"dateDeleted,omitempty" bson:"dateDeleted,omitempty"`
	ExpirationDate time.Time          `json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	OwnerID        string             `json:"ownerId" bson:"ownerId"`
}

// SetFileExpirationError is an error that indicates that there was a problem setting a file as expired
type SetFileExpirationError string

func (sfve SetFileExpirationError) Error() string {
	return string(sfve)
}

// ValidationError is an error that relays what about a file object made it invalid.
type ValidationError map[string]string

func (fve ValidationError) Error() string {
	var errorString string = "File object was invalid:\n"
	for k := range fve {
		errorString += fmt.Sprintf("\t%v: %v\n", k, fve[k])
	}
	return errorString
}

// NewFile returns a new File object
func NewFile(mimeType, name, ownerID, storageType string) *File {
	return &File{
		CreatedDate: time.Now(),
		MimeType:    mimeType,
		Name:        name,
		OwnerID:     ownerID,
		StorageType: storageType,
	}
}

// Validate is a function that validates a File struct to make sure it has valid data.
func (f *File) Validate() error {
	errorMessages := make(map[string]string)

	if f == nil {
		errorMessages["General"] = "the file cannot be nil."
		return ValidationError(errorMessages)
	}

	if f.MimeType == "" {
		errorMessages["MimeType"] = fmt.Sprint("MimeType cannot be empty")
	}

	if f.StorageType == "" {
		errorMessages["StorageType"] = fmt.Sprint("StorageType cannot be empty")
	}

	if f.OwnerID == "" {
		errorMessages["OwnerID"] = fmt.Sprint("OwnerID cannot be empty")
	}

	if f.Name == "" {
		errorMessages["Name"] = fmt.Sprint("Name cannot be empty")
	}

	if len(errorMessages) > 0 {
		return ValidationError(errorMessages)
	}

	return nil
}

// Delete sets the deleted date on a file.
func (f *File) Delete() {
	f.DeletedDate = time.Now()
}

// SetExpiration sets the expiration date on a file to the parameter
func (f *File) SetExpiration(expirationDate time.Time) error {
	if f.IsExpired() == true {
		return SetFileExpirationError(fmt.Sprintf("%v - %v is already expired: current expiration date %v", f.FileID, f.Name, f.ExpirationDate))
	}
	f.ExpirationDate = expirationDate
	return nil
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
