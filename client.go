package gosound

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Auth stores information needed to authorize
type Auth struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AccessToken  string
}

const (
	AuthEndpoint  = "https://soundcloud.com/connect"
	TokenEndpoint = "https://api.soundcloud.com/oauth2/token"
	APIURL        = "https://api.soundcloud.com/"
)

type Client struct {
	ID      string
	Secret  string
	AuthURL string
	client  *http.Client
}

func NewClient(auth Auth) (Client, error) {
	if auth.ClientID == "" {
		return Client{}, errors.New("required ClientID not given in parameter")
	}
	c := Client{
		ID:     auth.ClientID,
		Secret: auth.ClientSecret,
		client: &http.Client{},
	}

	v := url.Values{}
	v.Set("scope", "non-expiring")
	v.Set("client_id", auth.ClientID)
	v.Set("response_type", "code")
	v.Set("redirect_uri", auth.RedirectURI)

	c.AuthURL = AuthEndpoint + "?" + v.Encode()
	return c, nil
}

func (c *Client) Request(method, url string, params url.Values) []byte {
	params.Set("client_id", c.ID)
	reqUrl := APIURL + url + "?" + params.Encode()
	req, _ := http.NewRequest(method, reqUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, _ := c.client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func (c *Client) AddClient(url string) string {
    return url+"?client_id="+c.ID
}
