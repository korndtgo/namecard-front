package auth

import (
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

//JWTToken ...
type JWTClaims struct {
	*jwt.StandardClaims
	JWTPayload
}

//JWTPayload ...
type JWTPayload struct {
	Type          string      `json:"type"`
	Secret        interface{} `json:"secret"`
	PublicKeyPath string
}

type Auth struct {
	UserId   string `json:"user_id"`
	DeviceId string `json:"device_id"`
	IsAdmin  bool
}

//JWTTokenType ...
type JWTTokenType string

const (
	GraphQL JWTTokenType = "Graphql"
)

func (j *JWTPayload) ToPayload(token string) error {
	claims := jwt.MapClaims{}
	file, err := ioutil.ReadFile(j.PublicKeyPath)
	if err != nil {
		log.Println(j.PublicKeyPath)
		return err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(file)
	if err != nil {
		return err
	}

	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return err
	}

	j.Type = claims["type"].(string)
	j.Secret = claims["secret"].(string)

	return nil
}
