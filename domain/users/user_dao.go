package users

import (
	"fmt"
	"vanthanh.com/bookstore_users_api/datasources/mysql"
	"vanthanh.com/bookstore_users_api/utils/errors"
	"vanthanh.com/bookstore_users_api/utils/mysql_utils"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, created_at, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser      = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE id=?;"
	queryUpdateUser   = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser   = "DELETE FROM users WHERE id=?;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.ResErr {
	stmt, err := mysql.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.ResponseServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Update() *errors.ResErr {
	stmt, err := mysql.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.ResponseServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
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

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.Status, user.Password)
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

func (user *User) Delete() *errors.ResErr {
	stmt, err := mysql.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.ResponseServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
func (user *User) FindByStatus(status string) ([]User, *errors.ResErr) {
	stmt, err := mysql.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.ResponseServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.ResponseServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.ResponseNotFound(fmt.Sprintf("Not found user have status is %s", status))
	}
	return results, nil
}
