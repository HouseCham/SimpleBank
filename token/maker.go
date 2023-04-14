package token

import "time"

type Maker interface {
	// CreateToken creates a token for the given username and duration.
	CreateToken(username string, duration time.Duration) (string, error)
	
	// VerifyToken verifies the given token and returns the payload.
	VerifyToken(token string) (*Payload, error)
}