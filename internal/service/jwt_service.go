package service

import (
	"os"
	"time"
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type SecretJwk struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	K   string `json:"k"`
	Alg string `json:"alg"`
}

var JwtSignK []byte

func LoadJwtSignKeyInCache() error {

	file, err := os.Open("secret.jwk")

	if err != nil {
		logrus.Printf("Error while opening JWT secret key file: %s", err.Error())
		return err
	}

	var secretJwk SecretJwk

	dec := json.NewDecoder(file)
	err = dec.Decode(&secretJwk)

	if err != nil {
		logrus.Printf("Error while decoding JWT secret sign key: %s", err.Error())
		return err
	}

	logrus.Printf("%+v\n", secretJwk)

	JwtSignK = []byte(secretJwk.K)

	return nil
}

func ServeJwt(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    jwtToken, err := token.SignedString(JwtSignK)

	if err != nil {
		return "", err
	}

	logrus.Println("Generated Jwt token: ", jwtToken)
	
	return jwtToken, nil

}
