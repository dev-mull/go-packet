package packet

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	DefaultTimeout time.Duration = time.Second * 5
)

func NewAPI(config *Config) (*API, error) {
	//Setup defaults
	if config.Timeout == nil {
		config.Timeout = &DefaultTimeout
	}

	//Validate we have what we need
	if config.Token == "" {
		return nil,errors.New("API token not provided")
	}

	//Fire it up
	a := &API{
		Config: config,
		Url:    "https://api.packet.net", // so we can run httptest
	}
	return a,nil

}

func (a *API) httpRequest(method string, route string, body io.Reader) (*http.Response, error) {
	client := &http.Client{
		Timeout: *a.Config.Timeout,
	}

	req, err := http.NewRequest(method, a.Url+route, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Auth-Token", a.Config.Token)

	response, err := client.Do(req)
	if a.Debug {
		fmt.Printf("HTTP REQ. Method: %v URL: %v Code: %v\n", req.Method, req.URL, response.StatusCode)
	}
	return response, err

}

type APIErrors struct {
	Errors []string `json:"errors"`
}

type Config struct {
	Token     string        `yaml:"token"`
	ProjectID string        `yaml:"project-id"`
	Timeout   *time.Duration `yaml:"timeout"`
}

type API struct {
	Config *Config
	Debug  bool
	Url    string // so we can run http test
}

