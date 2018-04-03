package wow

import (
	"github.com/w0xel/go-bnet/internal"
	"net/http"
)

type WowService struct {
	client *internal.Client
	httpClient * http.Client
}

type ClassSlice struct {
	Classes []Class `json:"classes"`
}

func NewService(client * internal.Client) *WowService {
	return &WowService {
		client: client,
		httpClient: &http.Client {
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (wow * WowService) Classes() (*ClassSlice, *internal.Response, error) {
	req, err := wow.client.NewRequest("GET", "wow/data/character/classes?locale=en_GB", nil)
	if err != nil {
		return nil, nil, err
	}

	var classes ClassSlice
	resp, err := internal.Do(req, &classes, wow.httpClient)
	if err != nil {
		return nil, resp, err
	}

	return &classes, resp, nil
}
