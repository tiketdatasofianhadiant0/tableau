package tableau

import (
	"fmt"
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

func (a *authentication) getBearerToken() string {
	if a.accessToken == "" {
		return ""
	}

	return fmt.Sprintf(bearerAuthorization, a.accessToken)
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
//
//	POST /api/api-version/auth/signin
//
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_authentication.htm#sign_in
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
				ContentUrl: &a.base.cfg.ContentUrl,
			},
		},
	}

	url := a.base.cfg.GetUrl(signInUri)
	if url == "" {
		return ErrInvalidHost
	}

	res, err := a.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeJSON).
		SetBody(reqBody).
		Post(url)

	a.base.SetResponse(*res)
	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusOK {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	resBody := models.SignInBody{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return ErrFailedUnmarshalResponseBody
	}

	ts := time.Now()
	a.signInAt = &ts
	a.accessToken = resBody.Credentials.Token
	a.userID = *resBody.Credentials.User.ID
	a.siteID = *resBody.Credentials.Site.ID

	return nil
}

// SignOut Signs you out of the current session.
// This call invalidates the authentication token that is created by a call to Sign In.
//
// URI:
//
//	POST /api/api-version/auth/signout
//
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_authentication.htm#sign_out
func (a *authentication) SignOut() error {
	if !a.IsSignedIn() {
		return nil
	}

	url := a.base.cfg.GetUrl(signOutUri)
	if url == "" {
		return ErrInvalidHost
	}

	res, err := a.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeJSON).
		SetHeader(authorizationHeader, a.getBearerToken()).
		Post(url)

	a.base.SetResponse(*res)
	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusNoContent {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	a.signInAt = nil
	a.accessToken = ""
	a.userID = ""
	a.siteID = ""

	return nil
}

// SwitchSite Switches you onto another site without having to provide a username and password again.
// This method allows an authenticated user to switch sites that they have access to.
// Using the current authentication token, the method signs you in as a user on the site specified in the request payload.
// The method returns a new authentication token and invalidates the old one.
// You have the permissions of the user associated with the authorization token.
// By default, the token is good for 120 minutes.
//
// URI:
//
//	POST /api/api-version/auth/switchSite
//
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_authentication.htm#switch_site
func (a *authentication) SwitchSite(contentUrl string) error {
	if !a.IsSignedIn() {
		if err := a.SignIn(); err != nil {
			return err
		}
	}

	if a.base.cfg.ContentUrl == contentUrl {
		return nil
	}

	reqBody := models.SwitchSiteBody{
		Site: &models.Site{
			ContentUrl: &contentUrl,
		},
	}

	url := a.base.cfg.GetUrl(switchSiteUri)
	if url == "" {
		return ErrInvalidHost
	}

	res, err := a.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeJSON).
		SetHeader(authorizationHeader, a.getBearerToken()).
		SetBody(reqBody).
		Post(url)

	a.base.SetResponse(*res)
	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusOK {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	resBody := models.SignInBody{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return ErrFailedUnmarshalResponseBody
	}

	ts := time.Now()
	a.signInAt = &ts
	a.accessToken = resBody.Credentials.Token
	a.userID = *resBody.Credentials.User.ID
	a.siteID = *resBody.Credentials.Site.ID
	a.base.cfg.ContentUrl = contentUrl

	return nil
}
