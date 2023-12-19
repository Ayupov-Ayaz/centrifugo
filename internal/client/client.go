package client

import (
	"ayupov-ayaz/centrifugo/internal/services/security"
	"crypto/tls"
	"github.com/centrifugal/centrifuge-go"
	"log"
	"time"
)

func defaultConfig(cfg security.TokenGeneratorConfig) (*centrifuge.Config, error) {
	generateToken := security.GetToken(cfg)
	getTokenFunc := func(_ centrifuge.ConnectionTokenEvent) (string, error) {
		return generateToken()
	}

	resp := &centrifuge.Config{
		GetToken:           getTokenFunc,
		Data:               nil,
		CookieJar:          nil,
		Header:             nil,
		Name:               "app-go",
		Version:            "v1.0.0",
		NetDialContext:     nil,
		ReadTimeout:        5 * time.Second,  // by default
		WriteTimeout:       1 * time.Second,  // by default
		HandshakeTimeout:   1 * time.Second,  // by default
		MaxServerPingDelay: 10 * time.Second, // by default
		TLSConfig:          &tls.Config{},
		EnableCompression:  false,
	}

	return resp, nil
}

func setHooks(cli *centrifuge.Client) {
	cli.OnConnecting(func(e centrifuge.ConnectingEvent) {
		log.Printf("connecting: %d (%s)\n", e.Code, e.Reason)
	})

	cli.OnConnected(func(e centrifuge.ConnectedEvent) {
		log.Printf("connected: clientID=%s, version=%s\n", e.ClientID, e.Version)
	})

	cli.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		log.Printf("disconnected: %d (%s) \n", e.Code, e.Reason)
	})

	cli.OnError(func(e centrifuge.ErrorEvent) {
		log.Printf("error: %s\n", e.Error.Error())
	})
}
