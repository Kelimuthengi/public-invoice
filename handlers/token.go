package handlers

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	Username    string `json:"foo"`
	UserID      uint   `json:"userId"`
	Email       string `json:"email"`
	Phonenumber string `json:"phone"`
	Address     string `json:"address"`
	jwt.RegisteredClaims
}

// create the claims

func GenerateNewToken(j JwtToken) (string, error) {
	claims := JwtToken{
		Username:    j.Username,
		UserID:      j.UserID,
		Email:       j.Email,
		Phonenumber: j.Phonenumber,
		Address:     j.Address,
	}
	signInKey := []byte(os.Getenv("JWT_KEY"))
	fmt.Println("signInKey", signInKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signInKey)
	return ss, err
}
