package disk

import (
	"os"
	"path/filepath"

	"calvinechols.com/vault/filestore"
)

type fileStoreDiskRepo struct{}

// NewFileStoreDiskRepo returns a new instance of the file storage disk repo.
func NewFileStoreDiskRepo() filestore.Repo {
	return &fileStoreDiskRepo{}
}

func (r *fileStoreDiskRepo) GetFileHandle(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *fileStoreDiskRepo) ReadFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, fi.Size())
	_, err = f.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *fileStoreDiskRepo) CreateFile(name string, data []byte) error {
	filePath := filepath.Dir(name)
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return err
}
