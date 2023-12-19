package subscription

import (
	"ayupov-ayaz/centrifugo/internal/services/security"
	"fmt"
	"github.com/centrifugal/centrifuge-go"
	"log"
)

func DefaultConfig(cfg security.TokenGeneratorConfig) (*centrifuge.SubscriptionConfig, error) {
	tokenGenerator := security.GetTokenWithClaims(cfg)

	refreshToken := func(e centrifuge.SubscriptionTokenEvent) (string, error) {
		token, err := tokenGenerator(map[string]interface{}{
			"channel": e.Channel,
		})

		if err != nil {
			return "", fmt.Errorf("generate token failed: %w", err)
		}

		fmt.Printf("subscribe to '%s', token = %s", e.Channel, token)

		return token, nil
	}

	resp := &centrifuge.SubscriptionConfig{
		Token:       "", // empty
		GetToken:    refreshToken,
		Positioned:  false,
		Recoverable: false,
		JoinLeave:   false,
	}

	return resp, nil
}

func setHooks(sub *centrifuge.Subscription) {
	sub.OnSubscribing(func(e centrifuge.SubscribingEvent) {
		log.Printf("subscribe to '%d' : '%s' \n", e.Code, e.Reason)
	})

	sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
		log.Printf("Subscribed on channel %s : %v", sub.Channel, e.Recovered)
	})

	sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
		log.Printf("Unsubscribed from channel %s - %d (%s)",
			sub.Channel, e.Code, e.Reason)
	})

	sub.OnError(func(e centrifuge.SubscriptionErrorEvent) {
		log.Printf("Subscription error: %s", e.Error.Error())
	})

	sub.OnJoin(func(e centrifuge.JoinEvent) {
		log.Printf("Join: %s", e.ClientInfo.User)
	})

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		log.Printf("Publication: %s", string(e.Data))
	})
}

func New(cli *centrifuge.Client, cfg security.TokenGeneratorConfig, channel string,
) (*centrifuge.Subscription, error) {
	configs, err := DefaultConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("default config failed: %w", err)
	}

	sub, err := cli.NewSubscription(channel, *configs)
	if err != nil {
		return nil, fmt.Errorf("new subscription to %s failed: %w", channel, err)
	}

	setHooks(sub)

	if err = sub.Subscribe(); err != nil {
		return nil, fmt.Errorf("subscribe to %s failed: %w", channel, err)
	}

	return sub, nil
}
