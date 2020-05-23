package access

import "go.mongodb.org/mongo-driver/bson/primitive"

// Repo is the interface for interacting with access controls for files in the vault.
type Repo interface {
	// AddAccess adds an access record to the database.
	AddAccess(access *Access) (string, error)
	// GetAccess retreives access details from the database.
	GetAccessByAccessToken(accessToken string) (*Access, error)
	// AddLog
	AddLog(log *Log) (primitive.ObjectID, error)
	// GetLogsByAccessID returns the access logs for a given AccessID
	GetLogsByAccessToken(accessToken string) ([]Log, error)
	// GetLogsByFileID returns the access logs for a given FileID
	GetLogsByFileID(fileID primitive.ObjectID) ([]Log, error)
	// GetLogsByFileToken returns the access logs for a given FileToken
	GetLogsByFileToken(fileToken string) ([]Log, error)
}
