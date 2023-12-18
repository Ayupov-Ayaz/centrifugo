package security

import (
	"fmt"
	"github.com/centrifugal/centrifuge-go"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GetToken(cfg TokenGeneratorConfig) func(_ centrifuge.ConnectionTokenEvent) (string, error) {
	return func(_ centrifuge.ConnectionTokenEvent) (string, error) {
		claims := jwt.MapClaims{
			"sub": cfg.AppKey,
			"exp": time.Now().Add(cfg.Expiration).Unix(),
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
			SignedString([]byte(cfg.Secret))
		if err != nil {
			return "", fmt.Errorf("error creating jwt token: %w", err)
		}

		fmt.Println("token: ", token)

		return token, nil
	}
}
