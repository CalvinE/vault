package access

import "go.mongodb.org/mongo-driver/bson/primitive"

// Repo is the interface for interacting with access controls for files in the vault.
type Repo interface {
	// AddAccess adds an access record to the database.
	AddAccess(access *Access) (*primitive.ObjectID, error)
	// GetAccess returns an access record from the database given the accessID
	GetAccess(accessID *primitive.ObjectID) (*Access, error)
	// GetAllAccessesForFileID returns all access records for a given file id
	GetAllAccessesForFileID(fileID *primitive.ObjectID, ownerID string) ([]Access, error)
	// DisableAccess disables an access given the accessID
	DisableAccess(accessID *primitive.ObjectID) error
	// AddLog
	AddLog(log *Log) (*primitive.ObjectID, error)
	// GetLogsByAccessID returns the access logs for a given AccessID
	GetLogsByAccessID(accessID *primitive.ObjectID) ([]Log, error)
	// GetLogsByFileID returns the access logs for a given FileID
	GetLogsByFileID(fileID *primitive.ObjectID) ([]Log, error)
}
