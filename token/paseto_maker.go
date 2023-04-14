package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PaseToMaker struct {
	paseto *paseto.V2
	symmetricKey []byte
}

func NewPaseToMaker(symmetricKey []byte) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("symmetric key must be %d bytes long", chacha20poly1305.KeySize)
	}

	maker := &PaseToMaker{
		paseto: paseto.NewV2(),
		symmetricKey: symmetricKey,
	}
	return maker, nil
}

// CreateToken creates a token for the given username and duration.
func (maker *PaseToMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}
	
// VerifyToken verifies the given token and returns the payload.
func (maker *PaseToMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}