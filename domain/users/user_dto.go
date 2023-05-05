package users

import (
	"strings"
	"vanthanh.com/bookstore_users_api/utils/errors"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (user *User) Validate() *errors.ResErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.ResponseError("invalid email address")
	}
	return nil
}
