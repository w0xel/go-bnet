package profile

import (
	"github.com/w0xel/go-bnet/client"
)

// AccountService has account-related APIs. See Client.
type AccountService struct {
	client *client.Client
}

func NewAccountService(c *client.Client) *AccountService {
	return &AccountService{client: c}
}

// User represents the user information for a Battle.net account
type User struct {
	ID        int    `json:"id"`
	BattleTag string `json:"battletag"`
}

// User calls the /account/user endpoint. See Battle.net docs.
func (s *AccountService) User() (*User, *client.Response, error) {
	req, err := s.client.NewRequest("GET", "account/user", nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}

