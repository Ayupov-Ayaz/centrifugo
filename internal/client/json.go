package client

import (
	"ayupov-ayaz/centrifugo/internal/services/security"
	"fmt"
	"github.com/centrifugal/centrifuge-go"
)

func NewJsonClient(
	endpoint string, cfg security.TokenGeneratorConfig) (*centrifuge.Client, error) {
	config, err := defaultConfig(cfg)
	if err != nil {
		return nil, err
	}

	cli := centrifuge.NewJsonClient(endpoint, *config)
	setHooks(cli)

	if err := cli.Connect(); err != nil {
		return nil, fmt.Errorf("error connecting to centrifugo: %w", err)
	}

	return cli, nil
}
