package client

import (
	"ayupov-ayaz/centrifugo/config"
	"fmt"
	"github.com/centrifugal/centrifuge-go"
)

func NewJsonClient(
	endpoint string, cfg *config.Config) (*centrifuge.Client, error) {
	cli := centrifuge.NewJsonClient(endpoint, defaultConfig(cfg))

	setHooks(cli)

	if err := cli.Connect(); err != nil {
		return nil, fmt.Errorf("error connecting to centrifugo: %w", err)
	}

	return cli, nil
}
