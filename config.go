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

func (c *Config) initConfig() error {
	h := strings.TrimSpace(c.Host)
	u, err := url.Parse(h)
	if len(h) == 0 || err != nil {
		return ErrInvalidHost
	}
	c.Host = u.String()

	if c.Version == "" {
		c.Version = DefaultVersion
	}

	if c.Username == "" || c.Password == "" {
		return ErrInvalidUsernamePassword
	}

	return nil
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
