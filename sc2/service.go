package sc2

import (
	"fmt"
	"github.com/w0xel/go-bnet/internal"
)

type SC2Service struct {
	client *internal.Client
}

func NewService(client * internal.Client) *SC2Service {
	return &SC2Service {
		client: client,
	}
}

// TODO: Create type to hold information about an individual SC2 profile's ladders.
// TODO: Create function for calling /sc2/profile/:id/:region/:name/ladders

// Match represents information about a single Starcraft 2 match.
type Match struct {
	Map      string `json:"map"`
	Type     string	`json:"type"`
	Decision string `json:"decision"`
	Speed    string `json:"speed"`
	Date     int    `json:"date"`
}

// Matches is a structure for holding multiple Match instances
type Matches struct {
	Matches []Match `json:"matches"`
}

// Matches(:id, :realm, :name) calls the /sc2/profile/:id/:realm/:name/matches endpoint.
// This provides data about an individual SC2 profile's match history.
func (s *SC2Service) Matches(id int, realm int, name string) (*Matches, *internal.Response, error) {
	url := fmt.Sprintf("profile/%d/%d/%s/matches", id, realm, name)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var matches Matches
	resp, err := s.client.Do(req, &matches)
	if err != nil {
		return nil, resp, err
	}

	return &matches, resp, nil
}
