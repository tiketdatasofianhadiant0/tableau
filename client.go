package tableau

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	c              *resty.Client
	cfg            *Config
	Authentication *authentication
}

func NewClient(cfg Config) (*Client, error) {
	if err := cfg.initConfig(); err != nil {
		return nil, err
	}

	restClient := resty.New()
	restClient.JSONMarshal = json.Marshal
	restClient.JSONUnmarshal = json.Unmarshal

	client := &Client{
		c:   restClient,
		cfg: &cfg,
	}

	auth := &authentication{base: client}
	client.Authentication = auth

	return client, nil
}
