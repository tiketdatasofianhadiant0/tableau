package tableau

import "github.com/go-resty/resty"

type Client struct {
	c              *resty.Client
	cfg            *Config
	Authentication *authentication
	host           string
	version        string
}
