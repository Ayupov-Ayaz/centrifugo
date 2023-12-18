package run

import (
	"ayupov-ayaz/centrifugo/internal/client"
	"ayupov-ayaz/centrifugo/internal/services/security"
	"fmt"
	"github.com/centrifugal/centrifuge-go"
	"time"
)

const host = "ws://localhost:8000/connection/websocket"

func cfg() client.Config {
	return client.Config{
		TokenGeneratorConfig: security.TokenGeneratorConfig{
			AppKey:     "tommy",
			Secret:     "secret-key",
			Expiration: 1 * time.Hour,
		},
		Version: "v1.0.0",
	}
}

func generateToken(cfg centrifuge.Config) (string, error) {
	token, err := cfg.GetToken(centrifuge.ConnectionTokenEvent{})
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return token, nil
}

func Run() error {
	config := client.DefaultConfig(cfg())

	token, err := generateToken(config)
	if err != nil {
		return fmt.Errorf("error generating token: %w", err)
	}
	config.Token = token

	cli := client.NewJsonClient(host, config)

	if err := cli.Connect(); err != nil {
		return fmt.Errorf("error connecting to centrifugo: %w", err)
	}

	defer cli.Close()

	time.Sleep(2 * time.Minute)

	return nil
}
