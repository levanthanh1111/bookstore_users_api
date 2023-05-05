package users

import (
	"vanthanh.com/bookstore_users_api/datasources/mysql"
	"vanthanh.com/bookstore_users_api/utils/errors"
	"vanthanh.com/bookstore_users_api/utils/mysql_utils"
	"vanthanh.com/bookstore_users_api/utils/time"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?"
)

func (user *User) Get() *errors.ResErr {
	stmt, err := mysql.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.ResponseServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.ResErr {
	stmt, err := mysql.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.ResponseServerError(err.Error())
	}
	defer stmt.Close()

	user.CreatedAt = time.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.ResponseServerError(err.Error())
	}

	user.Id = userId
	return nil
}
