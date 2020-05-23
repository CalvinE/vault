package file

import "go.mongodb.org/mongo-driver/bson/primitive"

// Repo is the interface for access meta file information from the database
type Repo interface {
	// AddFile adds a files information to the database.
	AddFile(file *File) (primitive.ObjectID, error)
	// GetFile gets a files information from the database.
	GetFile(id primitive.ObjectID) (*File, error)
}
