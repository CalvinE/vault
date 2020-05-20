package access

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Access this is a type representing access to files in the vault
type Access struct {
	AccessID          string    `json:"accessId"`
	FileID            string    `json:"fileId"`
	CreatedDate       time.Time `json:"createdDate"`
	ExpirationDate    time.Time `json:"expirationDate,omitempty"`
	AllowAnonymous    bool      `json:"allowAnonymous"`
	AnonymousPassword string    `json:"anonymousPassword,omitempty"`
	AccessCount       int64     `json:"accessCount"`
}

type DBAccess struct {
	DbID   primitive.ObjectID `json:"_id,omitempty"`
	Access `bson:",inline"`
}
