package service

import (
	"errors"
	"os"

	"github.com/adamnasrudin03/library/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)


type AuthService interface {
	GenerateToken(userID uint64, userName string)  (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	IsDuplicateEmail(email string) bool
}

type authService struct {
	publisherReopisory repository.PublisherRepository
}

func NewAuthService(publisherReopisory repository.PublisherRepository) *authService {
	return &authService{publisherReopisory}
}

func getSecretKey() string {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "adamnasrudin_key"
	}

	return secretKey
}

func (s *authService) GenerateToken(userID uint64, userName string) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	payload["user_name"] = userName

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	secretKey := []byte(getSecretKey())
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



func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.publisherReopisory.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
