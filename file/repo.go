package file

// Repo is the interface for access meta file information from the database
type Repo interface {
	// AddFile adds a files information to the database.
	AddFile(file *File) (string, error)
	// GetFile gets a files information from the database.
	// GetFile(id primitive.ObjectID) (*File, error)
	// GetFileByFileToken gets a files information from the data store by its fle token
	GetFileByFileToken(fileToken string) (*File, error)
}
