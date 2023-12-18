package client

import (
	"ayupov-ayaz/centrifugo/internal/services/security"
	"crypto/tls"
	"fmt"
	"github.com/centrifugal/centrifuge-go"
	"time"
)

type Config struct {
	security.TokenGeneratorConfig
	Version string
}

func DefaultConfig(cfg Config) centrifuge.Config {
	getTokenFunc := security.GetToken(cfg.TokenGeneratorConfig)

	return centrifuge.Config{
		GetToken:           getTokenFunc,
		Data:               nil,
		CookieJar:          nil,
		Header:             nil,
		Name:               "app-go",
		Version:            cfg.Version,
		NetDialContext:     nil,
		ReadTimeout:        5 * time.Second,  // by default
		WriteTimeout:       1 * time.Second,  // by default
		HandshakeTimeout:   1 * time.Second,  // by default
		MaxServerPingDelay: 10 * time.Second, // by default
		TLSConfig:          &tls.Config{},
		EnableCompression:  false,
	}
}

func NewJsonClient(endpoint string, cfg centrifuge.Config) *centrifuge.Client {
	client := centrifuge.NewJsonClient(endpoint, cfg)

	client.OnConnecting(func(e centrifuge.ConnectingEvent) {
		fmt.Printf("connecting: %d (%s)\n", e.Code, e.Reason)
	})

	client.OnConnected(func(e centrifuge.ConnectedEvent) {
		fmt.Printf("connected - clientID=%s\n", e.ClientID)
	})

	client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		fmt.Printf("disconnected: %d (%s) \n", e.Code, e.Reason)
	})

	client.OnError(func(e centrifuge.ErrorEvent) {
		fmt.Printf("error: %s\n", e.Error.Error())
	})

	return client
}
