package service

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)


type AuthService interface {
	GenerateToken(userID uint64, userName string)  (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func getSecretKey() string {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey != "" {
		secretKey = "adamnasrudin_key"
	}
	key := []byte(secretKey)

	return string(key)
}

func (s *authService) GenerateToken(userID uint64, userName string) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	payload["user_name"] = userName

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	secretKey := getSecretKey()
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		secretKey := getSecretKey()

		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

