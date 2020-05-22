package access

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Access this is a type representing access to files in the vault
type Access struct {
	AccessID          string    `json:"accessId" bson:"accessId"`
	FileID            string    `json:"fileId" bson:"fileId"`
	Name              string    `json:"name,omitempty" bson:"name,omitempty"`
	DisabledDate      time.Time `json:"disabledDate,omitempty" bson:"disabledDate,omitempty"`
	CreatedDate       time.Time `json:"createdDate" bson:"createdDate"`
	ExpirationDate    time.Time `json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	AllowAnonymous    bool      `json:"allowAnonymous" bson:"allowAnonymous"`
	AnonymousPassword string    `json:"anonymousPassword,omitempty" bson:"anonymousPassword,omitempty"`
	AccessCount       int64     `json:"accessCount" bson:"accessCount"`
}

// DBAccess represents an access from the database
type DBAccess struct {
	DbID   primitive.ObjectID `json:"_id,omitempty"`
	Access `bson:",inline"`
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
func NewAccess() *Access {
	return &Access{
		AccessID:    uuid.NewV4().String(),
		CreatedDate: time.Now(),
		AccessCount: 0,
	}
}

// Validate is a function that validates an Access struct to make sure it has valid data.
func (a *Access) Validate() ValidationError {
	errorMessages := make(map[string]string)

	if a == nil {
		errorMessages["General"] = "the access cannot be nil."
		return ValidationError(errorMessages)
	}

	if a.AccessID == "" {
		errorMessages["AccessID"] = fmt.Sprint("AccessID cannot be empty")
	}

	if a.FileID == "" {
		errorMessages["FileID"] = fmt.Sprint("FileID cannot be empty")
	}

	if a.AccessCount < 0 {
		errorMessages["AccessCount"] = fmt.Sprint("AccessCount must be greater than or equal to 0")
	}

	return ValidationError(errorMessages)
}

// IsDisabled returns true if the access' DisabledDate is set.
func (a *Access) IsDisabled() bool {
	unsetTime := time.Time{}
	wasDeleted := a.DisabledDate != unsetTime
	return wasDeleted
}

// IsExpired returns true if the file has an expiration data and it is after the current time.
func (a *Access) IsExpired() bool {
	unsetTime := time.Time{}
	hasExpirationDate := a.ExpirationDate != unsetTime
	if hasExpirationDate == true {
		now := time.Now()
		isExpired := now.After(a.ExpirationDate)
		return isExpired
	}
	return hasExpirationDate
}
