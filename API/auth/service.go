package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	err = godotenv.Load()
	key = os.Getenv("SECRET_KEY")
)

type Service interface {
	GenerateToken(ID uint32) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(ID uint32) (string, error) {

	err := godotenv.Load()

	if err != nil {
		log.Println(err)
	}

	key := os.Getenv("SECRET_KEY")

	claim := jwt.MapClaims{
		"user_id": ID,
	}

	//generate token useing HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {

		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
