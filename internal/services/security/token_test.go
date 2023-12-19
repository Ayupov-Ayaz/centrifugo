package security

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_GetToken(t *testing.T) {
	cfg := TokenGeneratorConfig{
		AppKey:     "app",
		Secret:     "secret",
		Expiration: 1 * time.Hour,
	}

	token, err := GetToken(cfg)()
	require.NoError(t, err)
	require.NotEmpty(t, token)

	const expPrefix = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."
	require.Equal(t, expPrefix, token[:len(expPrefix)])
	require.True(t, len(token) > len(expPrefix))
	fmt.Println(token)
}
