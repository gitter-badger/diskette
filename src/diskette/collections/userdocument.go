package collections

import (
	"time"

	"labix.org/v2/mgo/bson"
)

const UserCollectionName = "user"

type UserDocument struct {
	Id               bson.ObjectId          `bson:"_id"`
	Name             string                 `bson:"name"`
	Email            string                 `bson:"email"`
	HashedPass       []byte                 `bson:"hashedPass"`
	Language         string                 `bson:"language"`
	Roles            []string               `bson:"roles"`
	CreatedAt        time.Time              `bson:"createdAt"`
	ConfirmationKey  string                 `bson:"confirmationKey"`
	ConfirmedAt      time.Time              `bson:"confirmedAt"`
	ResetKey         string                 `bson:"resetKey"`
	RequestedResetAt time.Time              `bson:"requestedResetAt"`
	IsSuspended      bool                   `bson:"isSuspended"`
	SignedOutAt      time.Time              `bson:"signedOutAt"`
	Profile          map[string]interface{} `bson:"profile"`
}
