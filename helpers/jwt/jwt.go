package jwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenMetadata ...
type TokenMetadata struct {
	UserID int64
	Email  string
}

// CreateToken ...
func CreateToken(userID int64, email string) (string, error) {
	// tokenExp, _ := time.ParseDuration(os.Getenv("TOKEN_EXP"))

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyToken(headerToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid ...
func TokenValid(headerToken string) error {
	token, err := verifyToken(headerToken)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata ...
func ExtractTokenMetadata(headerToken string) (*TokenMetadata, error) {
	token, err := verifyToken(headerToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &TokenMetadata{
			UserID: userID,
			Email:  email,
		}, nil
	}
	return nil, err
}
