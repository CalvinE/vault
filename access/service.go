package access

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service is a conduit to access use cases related to access and access log records.
type Service interface {
	AddAccess(access *Access) (*primitive.ObjectID, error)
	GetAccess(accessID *primitive.ObjectID) (*Access, error)
	ValidateAccess(accessID, password, userID string) (*primitive.ObjectID, error)
	AddAccessLog(accessLog *Log) (*primitive.ObjectID, error)
}

type service struct {
	DBRepo Repo
}

// ValidateAccessError is an error indicating that access validationhas failed.
type ValidateAccessError struct {
	AccessID string
	Message  string
}

func (vae ValidateAccessError) Error() string {
	return fmt.Sprintf("Validation of access %v failed with: %v", vae.AccessID, vae.Message)
}

// NewAccessService returns a new access service
func NewAccessService(accessRepo Repo) Service {
	return &service{
		DBRepo: accessRepo,
	}
}

func (s *service) AddAccess(access *Access) (*primitive.ObjectID, error) {
	accessID, err := s.DBRepo.AddAccess(access)
	if err != nil {
		return nil, err
	}
	return accessID, nil
}

func (s *service) GetAccess(accessID *primitive.ObjectID) (*Access, error) {
	access, err := s.DBRepo.GetAccess(accessID)
	if err != nil {
		return nil, err
	}
	return access, nil
}

func (s *service) ValidateAccess(accessIDString, password, userID string) (*primitive.ObjectID, error) {
	accessID, err := primitive.ObjectIDFromHex(accessIDString)
	if err != nil {
		return nil, err
	}
	access, err := s.GetAccess(&accessID)
	if err != nil {
		return nil, err
	}
	if access.IsDisabled() == true {
		// the access has been disabled
		return nil, ValidateAccessError{
			AccessID: accessIDString,
			Message:  "the access has been disabled",
		}
	}
	if access.IsExpired() == true {
		// this access has expired
		return nil, ValidateAccessError{
			AccessID: accessIDString,
			Message:  "this access has expired",
		}
	}
	// TODO: implement specific user validation
	// If no specific users specified, check for anonymous access.
	if access.AllowAnonymous == true {
		if access.AnonymousPassword != "" {
			// anonymous access with password is required.
			if password == access.AnonymousPassword {
				return &access.FileID, nil
			}
			// the anonymous password did not match
			return nil, ValidateAccessError{
				AccessID: accessIDString,
				Message:  "access password did not match",
			}
		}
		// anonymous access without password is allowed
		return &access.FileID, nil
	}
	return nil, ValidateAccessError{
		AccessID: accessIDString,
		Message:  fmt.Sprintf("provided user did not have access: %v", userID),
	}
}

func (s *service) AddAccessLog(accessLog *Log) (*primitive.ObjectID, error) {
	accessLogID, err := s.DBRepo.AddLog(accessLog)
	if err != nil {
		return nil, err
	}
	return accessLogID, err
}
