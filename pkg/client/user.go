package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// User is ...
type User struct {
	space  string
	apiKey string
	values url.Values
}

// NewUser is constructor.
func NewUser(space string, apiKey string) *User {
	u := new(User)
	u.space = space
	u.apiKey = apiKey
	u.values = url.Values{}

	return u
}

// List function returns list of users in your space.
func (u *User) List() ([]UserResponse, error) {
	api := "api/v2/users"

	cli := NewClient(u.space, u.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []UserResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// ListOfProject function returns list of project members.
func (u *User) ListOfProject(projectKey string) ([]UserResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/users", projectKey)

	cli := NewClient(u.space, u.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []UserResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// PrintCSV function prints list of users in CSV format.
func (u *User) PrintCSV(r []UserResponse) {
	for _, n := range r {
		fmt.Printf("%d,%s,%s,%d,%s,%s\n", n.ID, n.UserID, n.Name, n.RoleType, n.Lang, n.MailAddress)
	}
}
