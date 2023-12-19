package run

import (
	"ayupov-ayaz/centrifugo/internal/client"
	"ayupov-ayaz/centrifugo/internal/services/security"
	"ayupov-ayaz/centrifugo/internal/subscription"
	"time"
)

const host = "ws://localhost:8000/connection/websocket"

func cfg() security.TokenGeneratorConfig {
	return security.TokenGeneratorConfig{
		AppKey:     "tommy",
		Secret:     "secret-key",
		Expiration: 1 * time.Hour,
	}
}

func Run() error {
	configs := cfg()
	cli, err := client.NewJsonClient(host, configs)
	if err != nil {
		return err
	}

	defer cli.Close()

	_, err = subscription.New(cli, configs, "news")

	//fmt.Println(sub)
	time.Sleep(10 * time.Hour)

	return nil
}
