package access

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Access this is a type representing access to files in the vault
type Access struct {
	AccessID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FileToken         string             `json:"fileToken" bson:"fileToken"`
	AccessToken       string             `json:"accessToken" bson:"accessToken"`
	Name              string             `json:"name,omitempty" bson:"name,omitempty"`
	DisabledDate      time.Time          `json:"disabledDate,omitempty" bson:"disabledDate,omitempty"`
	CreatedDate       time.Time          `json:"createdDate" bson:"createdDate"`
	ExpirationDate    time.Time          `json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	AllowAnonymous    bool               `json:"allowAnonymous" bson:"allowAnonymous"`
	AnonymousPassword string             `json:"anonymousPassword,omitempty" bson:"anonymousPassword,omitempty"`
	AccessCount       int64              `json:"accessCount" bson:"accessCount"`
	CreatorID         string             `json:"creator" bson:"creator"`
}

// Log is a type that represents the attempted use of an Access to get a file, successful or not.
type Log struct {
	LogID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccessToken   string             `json:"accessId" bson:"accessId"`
	ClientIP      string             `json:"clientIp" bson:"clientIp"`
	FailureReason string             `json:"failureReason,omitempty" bson:"failureReason,omitempty"`
	AttemptDate   time.Time          `json:"attemptDate" bson:"attemptDate"`
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
func NewAccess(creatorID, fileToken string) *Access {
	return &Access{
		AccessCount: 0,
		AccessToken: uuid.NewV4().String(),
		CreatedDate: time.Now(),
		CreatorID:   creatorID,
		FileToken:   fileToken,
	}
}

// NewLog returns a new Log object
func NewLog(accessToken, clientIP, failureReason string) *Log {
	return &Log{
		AccessToken:   accessToken,
		ClientIP:      clientIP,
		FailureReason: failureReason,
		AttemptDate:   time.Now(),
	}
}

// Validate is a function that validates an Access struct to make sure it has valid data.
func (a *Access) Validate() error {
	errorMessages := make(map[string]string)

	if a == nil {
		errorMessages["General"] = "the access cannot be nil."
		return ValidationError(errorMessages)
	}

	if a.AccessToken == "" {
		errorMessages["AccessToken"] = fmt.Sprint("AccessToken cannot be empty")
	}

	if a.CreatorID == "" {
		errorMessages["CreatorID"] = fmt.Sprint("AccessToken cannot be empty")
	}

	if a.FileToken == "" {
		errorMessages["FileToken"] = fmt.Sprint("FileToken cannot be empty")
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
