package tableau

import (
	"github.com/tiketdatarisal/tableau/models"
	"net/http"
	"time"
)

type authentication struct {
	base        *Client
	signInAt    *time.Time
	accessToken string
	userID      string
	siteID      string
}

func (a *authentication) IsSignedIn() bool {
	if a.userID == "" || a.accessToken == "" || a.siteID == "" || a.signInAt == nil {
		return false
	}

	if time.Now().Sub(*a.signInAt) >= tokenLifetime {
		return false
	}

	return true
}

func (a *authentication) SignIn(force ...bool) error {
	forceSignIn := len(force) > 0 && force[0]
	if a.IsSignedIn() && !forceSignIn {
		return nil
	}

	reqBody := models.SignInBody{
		Credentials: &models.Credentials{
			Name:     a.base.cfg.Username,
			Password: a.base.cfg.Password,
			Site: &models.Site{
				ContentUrl: a.base.cfg.ContentUrl,
			},
		},
	}

	url := a.base.cfg.GetUrl(authPath)
	if url == "" {
		return ErrInvalidHost
	}

	res, err := a.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJson).
		SetHeader(acceptHeader, mimeTypeJson).
		SetBody(reqBody).
		Post(url)
	if err != nil {
		if errBody, err := models.NewErrorBody(res.Body()); err == nil {
			return errBody.Error
		}

		return err
	}

	if res.StatusCode() != http.StatusOK {
		if errBody, err := models.NewErrorBody(res.Body()); err == nil {
			return errBody.Error
		}

		return ErrUnknownError
	}

	resBody := models.SignInBody{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return ErrFailedUnmarshalResponseBody
	}

	ts := time.Now()
	a.signInAt = &ts
	a.accessToken = resBody.Credentials.Token
	a.userID = resBody.Credentials.User.ID
	a.siteID = resBody.Credentials.Site.ID
	
	return nil
}
