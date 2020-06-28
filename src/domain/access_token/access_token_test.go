package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	token := GetNewAccessToken()
	assert.False(t, token.isExpired(), "Brand new token should not be expired")
	assert.EqualValues(t, "", token.AccessToken, "new AccessToken should not have defined access token ")
	assert.True(t, token.UserId == 0, "new AccessToken should not have an associated user id")
	assert.True(t, token.ClientId == 0, "new AccessToken should not have an associated client id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	token := AccessToken{}
	assert.True(t, token.isExpired(), "Brand new Access token should be expired by Default")

	token.Expires = time.Now().UTC().Add(time.Hour * 3).Unix()
	assert.False(t, token.isExpired(), "Token expiring three hours from now should NOT be expired.")
}

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration Time should be 24 hours")
}
