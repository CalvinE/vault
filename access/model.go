package access

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Access this is a type representing access to files in the vault
type Access struct {
	AccessID          string    `json:"accessId" bson:"accessId"`
	FileID            string    `json:"fileId" bson:"fileId"`
	Name              string    `json:"name,omitempty" bson:"name,omitempty"`
	Disabled          bool      `json:"disabled" bson:"disabled"`
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

// NewAccess returns a new Access object
func NewAccess() *Access {
	return &Access{
		AccessID:    uuid.NewV4().String(),
		Disabled:    false,
		CreatedDate: time.Now(),
		AccessCount: 0,
	}
}
