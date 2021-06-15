package access

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Access this is a type representing access to files in the vault
type Access struct {
	AccessID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FileID            primitive.ObjectID `json:"fileId" bson:"fileId"`
	Name              string             `json:"name,omitempty" bson:"name,omitempty"`
	DisabledDate      *time.Time         `json:"disabledDate,omitempty" bson:"disabledDate,omitempty"`
	CreatedDate       *time.Time         `json:"createdDate" bson:"createdDate"`
	ExpirationDate    *time.Time         `json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	AllowAnonymous    bool               `json:"allowAnonymous" bson:"allowAnonymous"`
	AnonymousPassword string             `json:"anonymousPassword,omitempty" bson:"anonymousPassword,omitempty"`
	AccessCount       int64              `json:"accessCount" bson:"accessCount"`
	CreatorID         string             `json:"creator" bson:"creator"`
}

// Log is a type that represents the attempted use of an Access to get a file, successful or not.
type Log struct {
	LogID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccessID      primitive.ObjectID `json:"accessId" bson:"accessId"`
	FileID        primitive.ObjectID `json:"fileId" bson:"fileId"`
	ClientIP      string             `json:"clientIp" bson:"clientIp"`
	FailureReason string             `json:"failureReason,omitempty" bson:"failureReason,omitempty"`
	AttemptDate   *time.Time         `json:"attemptDate" bson:"attemptDate"`
}

// ValidationError is an error that relays what about a access object made it invalid.
type ValidationError map[string]string

func (fve ValidationError) Error() string {
	var errorString string = "Access object was invalid:\n"
	for k := range fve {
		errorString += fmt.Sprintf("\t%v: %v\n", k, fve[k])
	}
	return errorString
}

// NewAccess returns a new Access object
func NewAccess(fileID primitive.ObjectID, creatorID string) *Access {
	now := time.Now()
	return &Access{
		AccessCount: 0,
		// AccessToken: uuid.NewV4().String(),
		CreatedDate: &now,
		CreatorID:   creatorID,
		// FileToken:   fileToken,
		FileID: fileID,
	}
}

// NewLog returns a new Log object
func NewLog(accessID, fileID primitive.ObjectID, clientIP, failureReason string) *Log {
	now := time.Now()
	return &Log{
		AccessID:      accessID,
		FileID:        fileID,
		ClientIP:      clientIP,
		FailureReason: failureReason,
		AttemptDate:   &now,
	}
}

// Validate is a function that validates an Access struct to make sure it has valid data.
func (a *Access) Validate() error {
	errorMessages := make(map[string]string)
	emptyObjectID := primitive.ObjectID{}
	if a == nil {
		errorMessages["General"] = "the access cannot be nil."
		return ValidationError(errorMessages)
	}

	if a.CreatorID == "" {
		errorMessages["CreatorID"] = "AccessToken cannot be empty"
	}

	if a.FileID == emptyObjectID {
		errorMessages["FileID"] = "FileID cannot be empty"
	}

	if a.AccessCount < 0 {
		errorMessages["AccessCount"] = "AccessCount must be greater than or equal to 0"
	}

	return ValidationError(errorMessages)
}

// IsDisabled returns true if the access' DisabledDate is set.
func (a *Access) IsDisabled() bool {
	wasDeleted := a.DisabledDate != nil
	return wasDeleted
}

// IsExpired returns true if the file has an expiration data and it is after the current time.
func (a *Access) IsExpired() bool {
	hasExpirationDate := a.ExpirationDate != nil
	if hasExpirationDate {
		now := time.Now()
		isExpired := now.After(*a.ExpirationDate)
		return isExpired
	}
	return hasExpirationDate
}
