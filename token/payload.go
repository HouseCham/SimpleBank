package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = fmt.Errorf("invalid token")
	ErrExpiredToken = fmt.Errorf("token is expired")
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssueAt  time.Time `json:"issue_at"`  // time.Now()
	ExpireAt time.Time `json:"expire_at"` // time.Now().Add(duration)
}

// NewPayload creates a new payload for the given username and duration.
func NewPayload(username string, duration time.Duration) (*Payload, error) {

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:       tokenID,
		Username: username,
		IssueAt:  time.Now(),
		ExpireAt: time.Now().Add(duration),
	}, nil
}

// IsValid checks if the payload is valid.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpiredToken
	}
	return nil
}