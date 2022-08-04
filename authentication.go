package tableau

import "time"

type authentication struct {
	base        *Client
	signInAt    *time.Time
	accessToken string
	userID      string
	siteID      string
}
