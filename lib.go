package bnet

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"github.com/w0xel/go-bnet/internal"
	"github.com/w0xel/go-bnet/profile"
	"github.com/w0xel/go-bnet/wow"
)


const (
	libraryVersion = "0.1"
	userAgent      = "go-bnet/" + libraryVersion
)

type Client struct {
	internalClient internal.Client
}

// NewClient creates a new Battle.net client.
//
// region must be a valid Battle.net region. This will not validate it
// is valid.
//
// The http.Client argument should usually be retrieved via the
// oauth2 Go library NewClient function. It must be a client that
// automatically injects authentication details into requests.
func NewClient(region string, c *http.Client, apikey string) *Client {
	region = strings.ToLower(region)

	if c == nil {
		c = http.DefaultClient
	}

	// Determine the API base URL based on the region
	baseURLStr := fmt.Sprintf("https://%s.api.battle.net/", region)
	if region == "cn" {
		baseURLStr = "https://api.battlenet.com.cn/"
	}

	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		// We panic because we manually construct it above so it should
		// never really fail unless the user gives us a REALLY bad region.
		panic(err)
	}
	return &Client{
		internalClient: internal.Client {
			Client:		c,
			
			BaseURL:	baseURL,
			
			UserAgent:	userAgent,

			ApiKey: 	apikey,
		},
	}
}


// Hook to Account service.
func (c * Client) Account() *profile.AccountService {
	return profile.NewAccountService(&c.internalClient)
}

// Hook to Profile service.
func (c * Client) Profile() *profile.ProfileService {
	return profile.NewProfile(&c.internalClient)
}

func (c * Client) Wow() *wow.WowService {
	return wow.NewService(&c.internalClient)
}

func (c * Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	return c.internalClient.NewRequest(method, urlStr, body)
}

func (c * Client) Do(req *http.Request, v interface{}) (*internal.Response, error) {
	return c.internalClient.Do(req, v)
}
