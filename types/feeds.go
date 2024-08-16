package types

import (
	"time"
	"github.com/google/uuid"
)

type Feeds struct {
	ID         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"user_id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Url        string    `json:"url"`
}
