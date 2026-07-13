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

func ServeJwt(username string) (string, error) {

	logrus.Println("dev inprogress > Generate JWT with sign key: ")

	file, err := os.Open("secret.jwk")

	if err != nil {
		logrus.Printf("Error while opening JWT secret key file: %s", err.Error())
		panic(err)
	}

	var secretJwk SecretJwk

	dec := json.NewDecoder(file)
	_ = dec.Decode(&secretJwk)

	logrus.Printf("%+v\n", secretJwk)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    jwtToken, err := token.SignedString([]byte(secretJwk.K))

	if err != nil {
		return "", err
	}

	logrus.Println("Generated Jwt token: ", jwtToken)
	
	return jwtToken, nil

}
