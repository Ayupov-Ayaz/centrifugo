package run

import (
	"ayupov-ayaz/centrifugo/config"
	"ayupov-ayaz/centrifugo/internal/client"
	"ayupov-ayaz/centrifugo/internal/subscription"
	"fmt"
	"time"
)

const host = "ws://localhost:8000/connection/websocket"

func Run() error {
	configs, err := config.New()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

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
