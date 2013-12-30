package gosound

import (
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	EmptyAuth := Auth{}
	ClientAuth := Auth{ClientID: "c0a1760caac218e603dda520cc2966bf"}
	FullAuth := Auth{
		ClientID:     "c0a1760caac218e603dda520cc2966bf",
		ClientSecret: "41477ed2b037259706422d2aae466b82",
		RedirectURI:  "http://example.com/test",
	}
	FullAuth = FullAuth

	if _, err := NewClient(EmptyAuth); err == nil {
		t.Error("NewClient did not return an error with empty authentication.")
	}

	c, err := NewClient(ClientAuth)
	if err != nil {
		t.Error("NewClient returned with error with ClientID given")
	}

	_ = string(c.request("GET", "tracks", url.Values{}))
}
