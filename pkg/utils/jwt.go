package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

//   tao tocken cho account  cho nguoi dung dang nhap vao

func GenerateJWTTokenAuth(email string, uuid interface{}, role interface{}) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return "", err
	}
	secret := os.Getenv("SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["uuid"] = uuid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	//   lay id nguoi dung va cho vao set cua context de sau xac nhan cap q quyen truy cap cua ho

	return tokenString, nil

}

// giai ma token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	secret := os.Getenv("SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")

		}
		return []byte(secret), nil
	})

}
