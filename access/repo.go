package access

// Repo is the interface for interacting with access controls for files in the vault.
type Repo interface {
	// AddAccess adds an access record to the database.
	AddAccess(access *Access) (string, error)
	// GetAccess retreives access details from the database.
	GetAccess(id string) (*Access, error)
}
