package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"strings"
	"vanthanh.com/bookstore_users_api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.ResErr {
	getErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.ResponseError("Not found record")
		}
		return errors.ResponseServerError(err.Error())
	}
	switch getErr.Number {
	case 1062:
		return errors.ResponseError("invalid data")
	}
	return errors.ResponseServerError(err.Error())
}
