package security

import "time"

type TokenGeneratorConfig struct {
	AppKey     string
	Secret     string
	Expiration time.Duration
}
