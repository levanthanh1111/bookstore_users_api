package services

import (
	"vanthanh.com/bookstore_users_api/domain/users"
	"vanthanh.com/bookstore_users_api/utils/errors"
	"vanthanh.com/bookstore_users_api/utils/time"
)

func GetUser(userId int64) (*users.User, *errors.ResErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.ResErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func CreateUser(user users.User) (*users.User, *errors.ResErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.CreatedAt = time.GetNowDBString()
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId int64) *errors.ResErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func FindByStatus(status string) ([]users.User, *errors.ResErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
