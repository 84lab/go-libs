package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrParseWithClaims = errors.New("couldn't parse claims")
	ErrTokenExpired    = errors.New("token has been expired")
)

type Claim struct {
	Payload interface{} `json:"payload"`

	jwt.StandardClaims
}

func NewClaim(padyload interface{}, ttl time.Duration) *Claim {
	expiresAt := time.Now().Add(ttl)

	claims := &Claim{
		Payload:        padyload,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiresAt.Unix()},
	}

	return claims
}

func CreateToken(claims *Claim, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(signedToken, secret string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt token with claims: %w", err)
	}

	claims, ok := token.Claims.(*Claim)
	if !ok {
		return nil, ErrParseWithClaims
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, ErrTokenExpired
	}

	return claims, nil
}
