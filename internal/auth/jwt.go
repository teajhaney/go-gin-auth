package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)




type Claims struct{
	jwt.RegisteredClaims
	
	Role string `json:"role"`
}


func CreateToken(jwtSecret string, userID string, role string) (string, error){

	now := time.Now().UTC()	
	exp := now.Add(7*24 * time.Hour)

	claims := Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userID,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err :=	token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}
	return signedToken, nil
}



