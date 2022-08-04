package tableau

type usersGroups struct {
	base *Client
}

// AddUserToGroup Adds a user to the specified group.
//
// URI:
//   POST /api/api-version/sites/site-id/groups/group-id/users
func (u *usersGroups) AddUserToGroup(userID, groupID string) error {
	return nil
}
