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



// parse token 
func ParseToken(jwtSecret string, tokenString string) (Claims, error){
	var claims Claims
 parsedToken, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error){
	if token.Method.Alg()!= jwt.SigningMethodHS256.Alg(){
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(jwtSecret), nil
 }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
 if err != nil {
	return Claims{}, fmt.Errorf("error parsing token: %v", err)
 }
 if !parsedToken.Valid{
	return Claims{}, fmt.Errorf("invalid token")
 }

 if claims.Subject == ""{
	return Claims{}, fmt.Errorf("token missing suybject")
 }
 return claims, nil
}
