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

// SignIn Signs you in as a user on the specified site on Tableau Server or Tableau Online.
// This call returns a credentials token that you use in subsequent calls to the server.
// Typically, a credentials token is valid for 120 minutes.
//
// URI:
//   POST /api/api-version/auth/signin
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

	url := a.base.cfg.GetUrl(signInPath)
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

// SignOut Signs you out of the current session.
// This call invalidates the authentication token that is created by a call to Sign In.
//
// URI:
//   POST /api/api-version/auth/signout
func (a *authentication) SignOut() error {
	if !a.IsSignedIn() {
		return nil
	}

	url := a.base.cfg.GetUrl(signOutPath)
	if url == "" {
		return ErrInvalidHost
	}

	res, err := a.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJson).
		SetHeader(acceptHeader, mimeTypeJson).
		SetHeader(authorizationHeader, a.accessToken).
		Post(url)
	if err != nil {
		if errBody, err := models.NewErrorBody(res.Body()); err == nil {
			return errBody.Error
		}

		return err
	}

	if res.StatusCode() != http.StatusNoContent {
		if errBody, err := models.NewErrorBody(res.Body()); err == nil {
			return errBody.Error
		}

		return ErrUnknownError
	}

	a.signInAt = nil
	a.accessToken = ""
	a.userID = ""
	a.siteID = ""

	return nil
}
