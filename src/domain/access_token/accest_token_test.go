package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}

func TextGetNewAccentToken(t *testing.T) {
	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("Brand new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("New access token should not have defined acces token id")
	}

	if at.UserId != 0 {
		t.Error("New access token should not have an associated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("Empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	if at.IsExpired() {
		t.Error("Accces token expiring three hours from now should not be expired")
	}
}
