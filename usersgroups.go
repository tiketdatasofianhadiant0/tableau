package tableau

import (
	"fmt"
	"github.com/tiketdatarisal/tableau/models"
	"net/http"
	. "net/url"
	"strings"
)

type usersGroups struct {
	base *Client
}

// AddUserToGroup Adds a user to the specified group.
//
// URI:
//   POST /api/api-version/sites/site-id/groups/group-id/users
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#add_user_to_group
func (u *usersGroups) AddUserToGroup(userID, groupID string) (*models.User, error) {
	if !u.base.Authentication.IsSignedIn() {
		if err := u.base.Authentication.SignIn(); err != nil {
			return nil, err
		}
	}

	reqBody := models.UserBody{
		User: &models.User{
			ID: userID,
		},
	}

	url := u.base.cfg.GetUrl(fmt.Sprintf(addUserToGroupPath, u.base.Authentication.siteID, groupID))
	if url == "" {
		return nil, ErrInvalidHost
	}

	res, err := u.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJson).
		SetHeader(acceptHeader, mimeTypeJson).
		SetHeader(authorizationHeader, u.base.Authentication.getBearerToken()).
		SetBody(reqBody).
		Post(url)
	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusOK {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	resBody := models.UserBody{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return nil, ErrFailedUnmarshalResponseBody
	}

	return resBody.User, nil
}

// AddUserToSite Adds a user to Tableau Server or Tableau and assigns the user to the specified site.
// If Tableau Server is configured to use local authentication, the information you specify is used to create a new user in Tableau Server.
// When you add user to Tableau Online, the name of the user must be the email address that is used to sign in to Tableau Online.
// After you add a user, Tableau Online sends the user an email invitation.
// The user can click the link in the invitation to sign in and update their full name and password.
//
// URI:
//   POST /api/api-version/sites/site-id/users
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#add_user_to_site
func (u *usersGroups) AddUserToSite(user *models.User) (*models.User, error) {
	if !u.base.Authentication.IsSignedIn() {
		if err := u.base.Authentication.SignIn(); err != nil {
			return nil, err
		}
	}

	reqBody := models.UserBody{
		User: &models.User{
			Name:     user.Name,
			SiteRole: user.SiteRole,
		},
	}

	url := u.base.cfg.GetUrl(fmt.Sprintf(addUserToSitePath, u.base.Authentication.siteID))
	if url == "" {
		return nil, ErrInvalidHost
	}

	res, err := u.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJson).
		SetHeader(acceptHeader, mimeTypeJson).
		SetHeader(authorizationHeader, u.base.Authentication.getBearerToken()).
		SetBody(reqBody).
		Post(url)
	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusOK {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	resBody := models.UserBody{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return nil, ErrFailedUnmarshalResponseBody
	}

	return resBody.User, nil
}

// CreateGroup Creates a group.
//
// URI:
//   POST /api/api-version/sites/site-id/groups
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#create_group
func (u *usersGroups) CreateGroup(group *models.Group) (*models.Group, error) {
	return nil, nil
}

// DeleteGroup Deletes the group on a specific site.
// Deleting a group does not delete the users in group, but users are no longer members of the group.
// Any permissions that were previously assigned to the group no longer apply.
//
// URI:
//   DELETE /api/api-version/sites/site-id/groups/group-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#delete_group
func (u *usersGroups) DeleteGroup(groupID string) error {
	return nil
}

// GetGroupsForUser Gets a list of groups of which the specified user is a member.
//
// URI:
//   GET /api/api-version/sites/site-id/users/user-id/groups
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#get_groups_for_a_user
func (u *usersGroups) GetGroupsForUser(userID string, groupNames ...string) ([]models.Group, error) {
	if !u.base.Authentication.IsSignedIn() {
		if err := u.base.Authentication.SignIn(); err != nil {
			return nil, err
		}
	}

	query := ""
	if len(groupNames) >= 1 {
		query = fmt.Sprintf(
			"&filter=name:in:[%s]",
			strings.Replace(
				QueryEscape(strings.Join(groupNames, ",")),
				"%2C",
				",",
				-1))
	}

	pageNum := 1
	var result []models.Group
	for {
		url := u.base.cfg.GetUrl(fmt.Sprintf(getGroupsForUserPath, u.base.Authentication.siteID, userID))
		if url == "" {
			return nil, ErrInvalidHost
		}

		url = fmt.Sprintf(pagingParams, url, pageSize, pageNum, query)
		res, err := u.base.c.R().
			SetHeader(contentTypeHeader, mimeTypeJson).
			SetHeader(acceptHeader, mimeTypeJson).
			SetHeader(authorizationHeader, u.base.Authentication.getBearerToken()).
			Get(url)
		if err != nil {
			errBody, err := models.NewErrorBody(res.Body())
			if err != nil {
				return nil, ErrUnknownError
			}

			return nil, errCodeMap[errBody.Error.Code]
		}

		if res.StatusCode() != http.StatusOK {
			errBody, err := models.NewErrorBody(res.Body())
			if err != nil {
				return nil, ErrUnknownError
			}

			return nil, errCodeMap[errBody.Error.Code]
		}

		resBody := models.QueryGroupBody{}
		if err = json.Unmarshal(res.Body(), &resBody); err != nil {
			return nil, ErrFailedUnmarshalResponseBody
		}

		result = append(result, resBody.Groups.Group...)
		if pageNum*pageSize >= resBody.Pagination.GetTotalAvailable() {
			break
		}

		pageNum++
	}

	return result, nil
}

// GetUsersInGroup Gets a list of users in the specified group.
//
// URI:
//   GET /api/api-version/sites/site-id/groups/group-id/users
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#get_users_in_group
func (u *usersGroups) GetUsersInGroup(groupID string, userName ...string) ([]models.User, error) {
	return nil, nil
}

// GetUsersOnSite Returns the users associated with the specified site.
//
// URI:
//   GET /api/api-version/sites/site-id/users
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#get_users_on_site
func (u *usersGroups) GetUsersOnSite(userNames ...string) ([]models.User, error) {
	return nil, nil
}

// QueryGroups Returns a list of groups on the specified site, with optional parameters for specifying the paging of large results.
//
// URI:
//   GET /api/api-version/sites/site-id/groups
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#query_groups
func (u *usersGroups) QueryGroups(groupNames ...string) ([]models.Group, error) {
	return nil, nil
}

// QueryUserOnSite Returns information about the specified user.
//
// URI:
//   GET /api/api-version/sites/site-id/users/user-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#query_user_on_site
func (u *usersGroups) QueryUserOnSite() (*models.User, error) {
	return nil, nil
}

// RemoveUserFromSite Removes a user from the specified site.
// The user will be deleted if they do not own any other assets other than subscriptions.
// If a user still owns content (assets) on Tableau Server, the user cannot be deleted unless the ownership is reassigned first.
//
// You can’t remove a user from the server if they own content on any site on that server.
// You can remove a user from a site if they no longer own content on the site.
//
// If a user is removed from all sites that the user is a member of, the user is deleted.
//
// URI:
//   DELETE /api/api-version/sites/site-id/users/user-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#remove_user_from_site
func (u *usersGroups) RemoveUserFromSite(userID string, newUserID ...string) error {
	return nil
}

// RemoveUserFromGroup Removes a user from the specified group.
//
// URI:
//   DELETE /api/api-version/sites/site-id/groups/group-id/users/user-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#remove_user_to_group
func (u *usersGroups) RemoveUserFromGroup(userID, groupID string) error {
	return nil
}

// UpdateGroup Updates a group.
// If Tableau Server or Tableau Online site is configured to use local authentication, the method lets you update the group name.
//
// URI:
//   PUT /api/api-version/sites/site-id/groups/group-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#update_group
func (u *usersGroups) UpdateGroup(group *models.Group) (*models.Group, error) {
	return nil, nil
}

// UpdateUser Modifies information about the specified user.
// If Tableau Server is configured to use local authentication, you can update the user's name, email address, password, or site role.
// For Tableau Online, you can update the site role for a user, but you cannot update or change a user's password, user name (email address), or full name.
//
// URI:
//   PUT /api/api-version/sites/site-id/users/user-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm#update_user
func (u *usersGroups) UpdateUser(user *models.User) (*models.User, error) {
	return nil, nil
}
