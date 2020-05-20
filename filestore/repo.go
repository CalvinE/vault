package filestore

import "os"

// Repo is a generic interface for any code that access the actual files from the file storage system.
type Repo interface {
	// GetFileHandle get a handle to a file
	GetFileHandle(name string) (*os.File, error)
	// ReafFile returns the raw bytes of a file
	ReadFile(name string) ([]byte, error)
	// CreateFile creates a file given the name and the data as a byte array.
	CreateFile(name string, data []byte) error
}
