package models

import (
	"fmt"
)

const (
	apiUserByNickname = "users/%s"
)

// userData is a type that represents the way user is presented
// on the third party API server.
type userData struct {
	Data *User `json:"user"`
}

// User is a type that represents a single customer
// who buyes a product.
type User struct {
	Nickname string `json:"username"`
	Email    string `json:"email"`
}

// UserByNickname gets a username and returns an associated user
// if he/she does exist. It returns a non-nil error otherwise
// or if a connection to the API cannot be established.
func UserByNickname(nickname string) (*User, error) {
	var d userData
	err := objectFromURN(fmt.Sprintf(apiUserByNickname, nickname), &d)
	if err != nil {
		return nil, err
	}
	return d.Data, nil
}
