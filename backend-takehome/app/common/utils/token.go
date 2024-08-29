package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secretKey = []byte("secret-key")

func CreateToken(id int, email string, tokenExpTime int64) (string, error) {
	tm := time.Now()
	tmAfterAdd := tm.Add(time.Duration(tokenExpTime) * time.Second)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"exp":   tmAfterAdd.Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	var id string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = fmt.Sprintf("%v", claims["id"].(float64))
	} else {
		fmt.Println(err)
	}

	return id, nil
}
