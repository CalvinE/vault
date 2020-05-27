package file

import (
	"errors"
	"os"

	"calvinechols.com/vault/filestore"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service is an interface that represents use cases for working with files in the vault
type Service interface {
	SaveFileToStorage(data []byte) (string, error)
	AddFile(file *File) (primitive.ObjectID, error)
	GetFileFromStorage(name string) ([]byte, error)
	GetFileHandleFromStorage(name string) (*os.File, error)
	GetFile(fileID primitive.ObjectID) (*File, error)
	RetreiveFile(fileIDString, ownerID string) (*File, *os.File, error)
}

type service struct {
	DBRepo      Repo
	StorageRepo filestore.Repo
}

// NewFileService returns a new instance of the file service.
func NewFileService(fileRepo Repo, storageRepo filestore.Repo) Service {
	return &service{
		DBRepo:      fileRepo,
		StorageRepo: storageRepo,
	}
}

func (s *service) SaveFileToStorage(data []byte) (string, error) {
	newFileName := uuid.NewV4().String()
	err := s.StorageRepo.CreateFile(newFileName, data)
	if err != nil {
		return "", err
	}
	return newFileName, nil
}

func (s *service) AddFile(file *File) (primitive.ObjectID, error) {
	err := file.Validate()
	if err != nil {
		return primitive.NilObjectID, err
	}
	newFileID, err := s.DBRepo.AddFile(file)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return newFileID, nil
}

func (s *service) GetFileFromStorage(name string) ([]byte, error) {
	data, err := s.StorageRepo.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *service) GetFileHandleFromStorage(name string) (*os.File, error) {
	fileHandle, err := s.StorageRepo.GetFileHandle(name)
	if err != nil {
		return nil, err
	}
	return fileHandle, nil
}

func (s *service) GetFile(fileID primitive.ObjectID) (*File, error) {
	file, err := s.DBRepo.GetFile(fileID)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *service) RetreiveFile(fileIDString, ownerID string) (*File, *os.File, error) {
	fileID, err := primitive.ObjectIDFromHex(fileIDString)
	if err != nil {
		return nil, nil, err
	}
	f, err := s.GetFile(fileID)
	if err != nil {
		return nil, nil, err
	}
	if f.OwnerID == ownerID {
		fileHandler, err := s.GetFileHandleFromStorage(f.InternalFileName)
		if err != nil {
			return nil, nil, err
		}
		return f, fileHandler, nil
	}
	return nil, nil, errors.New("unauthorized attempt")
}
