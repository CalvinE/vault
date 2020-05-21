package file

// Repo is the interface for access meta file information from the database
type Repo interface {
	// AddFile adds a files information to the database.
	AddFile(file *File) (string, error)
	// GetFile gets a files information from the database.
	GetFile(id string) (*File, error)
}
