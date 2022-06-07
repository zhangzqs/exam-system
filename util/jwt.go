package util

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Jwt struct {
	secretKey       string
	expiresDuration time.Duration
}

func NewJwt(secretKey string, expiresDuration time.Duration) *Jwt {
	return &Jwt{
		secretKey:       secretKey,
		expiresDuration: expiresDuration,
	}
}

type customClaims struct {
	jwt.StandardClaims
	Uid int `json:"uid"`
}

func (j *Jwt) GenerateToken(uid int) string {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		customClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(j.expiresDuration).Unix(),
			},
			uid,
		},
	)
	token, _ := t.SignedString([]byte(j.secretKey))
	return token
}

func (j *Jwt) ParseToken(token string) (*customClaims, error) {
	var cc customClaims
	_, err := jwt.ParseWithClaims(token, &cc, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return &cc, err
	}
	return &cc, nil
}
