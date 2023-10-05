package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

// CreateToken creates a new token for a specific username and duration
// nel payload del token andremo a codificare l'username ed altre informazioni come l'expiration date!
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {

	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	// creiamo un nuovo token, passandogli il payload come claims
	// NB. claims e' un interface con un solo metodo Valid(..)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.secretKey))

}

func (o *JWTMaker) VerifyToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// ora devo prendere l'algoritmo... dall'header?
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			// se l'algoritmo trovato nell'header e' diverso --> dobbiamo ritornare errore
			return nil, ErrInvalidToken
		}
		return []byte(o.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
