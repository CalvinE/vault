package disk

import (
	"os"
	"path"
	"path/filepath"

	"calvinechols.com/vault/filestore"
)

type fileStoreDiskRepo struct {
	savePath string
}

// NewFileStoreDiskRepo returns a new instance of the file storage disk repo.
func NewFileStoreDiskRepo(savePath string) filestore.Repo {
	return &fileStoreDiskRepo{
		savePath: savePath,
	}
}

func (r *fileStoreDiskRepo) GetFileHandle(name string) (*os.File, error) {
	name = path.Join(r.savePath, name)
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *fileStoreDiskRepo) ReadFile(name string) ([]byte, error) {
	name = path.Join(r.savePath, name)
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
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
	name = path.Join(r.savePath, name)
	filePath := filepath.Dir(name)
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return err
}
