package jwt

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrParseWithClaims = errors.New("couldn't parse claims")
	ErrTokenExpired    = errors.New("token has been expired")
)

type Claim struct {
	UserID       string `json:"user_id"`
	RandomFactor int64  `json:"random_factor"` // allows to generate uniq jwt tokens within 1 sec
	jwt.StandardClaims
}

func NewClaim(userID string, ttl time.Duration) *Claim {
	expiresAt := time.Now().Add(ttl)

	randomNum, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		randomNum = &big.Int{}
	}

	claims := &Claim{
		UserID:         userID,
		RandomFactor:   randomNum.Int64(),
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
