package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func GetTokenWithClaims(cfg TokenGeneratorConfig) func(claims map[string]interface{}) (string, error) {
	return func(userClaims map[string]interface{}) (string, error) {
		claims := jwt.MapClaims{
			"sub": cfg.AppKey,
			"exp": time.Now().Add(cfg.Expiration).Unix(),
		}

		for k, v := range userClaims {
			claims[k] = v
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
			SignedString([]byte(cfg.Secret))
		if err != nil {
			return "", fmt.Errorf("error creating jwt token: %w", err)
		}

		log.Println("token: ", token)

		return token, nil
	}
}

func GetToken(cfg TokenGeneratorConfig) func() (string, error) {
	getToken := GetTokenWithClaims(cfg)

	return func() (string, error) {
		return getToken(nil)
	}
}
