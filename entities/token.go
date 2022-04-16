package entities

import (
	"log"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/errs"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Id    string
	Email string
	jwt.StandardClaims
}

func (t *Token) CreateAccessTokenString(signingMethod jwt.SigningMethod, secretKey string) (string, *errs.AppError) {

	tokenWithClaims := jwt.NewWithClaims(signingMethod, t)

	tokenString, err := tokenWithClaims.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("Error in generating token")
		return "", errs.NewUnexpectedError("Error in generating token")
	}
	return tokenString, nil
}
