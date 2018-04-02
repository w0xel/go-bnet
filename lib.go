package bnet

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"github.com/w0xel/go-bnet/client"
	"github.com/w0xel/go-bnet/profile"
)


const (
	libraryVersion = "0.1"
	userAgent      = "go-bnet/" + libraryVersion
)

type Client client.Client;

// NewClient creates a new Battle.net client.
//
// region must be a valid Battle.net region. This will not validate it
// is valid.
//
// The http.Client argument should usually be retrieved via the
// oauth2 Go library NewClient function. It must be a client that
// automatically injects authentication details into requests.
func NewClient(region string, c *http.Client) *Client {
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
		Client:    c,

		BaseURL:   baseURL,

		UserAgent: userAgent,
	}
}


// Hook to Account service.
func (c Client) Account() *profile.AccountService {
	cClient := client.Client(c)
	return profile.NewAccountService(&cClient)
}

// Hook to Profile service.
func (c Client) Profile() *profile.ProfileService {
	cClient := client.Client(c)
	return profile.NewProfile(&cClient)
}

func (c Client) BnetClient() *client.Client {
	cClient := client.Client(c)
	return &cClient
}
