package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

type JwtClaim struct {
	Email string
	jwt.RegisteredClaims
}

func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(j.ExpirationMinutes))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JwtClaim{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*JwtClaim); ok {
		if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
			return nil, errors.New("the session has expired, please login again")
		} else {
			return claims, nil
		}
	} else {
		return nil, errors.New("unknown claim type, cannot proceed")
	}
}