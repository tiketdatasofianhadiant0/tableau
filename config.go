package tableau

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

const (
	DefaultVersion = "3.10"
)

type Config struct {
	Host       string
	Version    string
	Username   string
	Password   string
	ContentUrl string
}

func initConfig(cfg Config) (*Config, error) {
	h := strings.TrimSpace(cfg.Host)
	u, err := url.Parse(h)
	if len(h) == 0 || err != nil {
		return nil, ErrInvalidHost
	}
	cfg.Host = u.String()

	if cfg.Version == "" {
		cfg.Version = DefaultVersion
	}

	if cfg.Username == "" || cfg.Password == "" {
		return nil, ErrInvalidUsernamePassword
	}

	return &cfg, nil
}

func (c *Config) GetUrl(paths ...string) string {
	u, err := url.Parse(c.Host)
	if err != nil {
		return ""
	}

	var ps []string
	ps = append(ps, u.Path)
	ps = append(ps, fmt.Sprintf("/api/%s", c.Version))
	ps = append(ps, paths...)
	u.Path = path.Join(ps...)

	return u.String()
}
