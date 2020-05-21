package file

import (
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
	WasDeleted     bool      `json:"deleted" bson:"deleted"`
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
		WasDeleted:  false,
	}
}
