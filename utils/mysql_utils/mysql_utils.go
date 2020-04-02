package mysql_utils

import (
	"github.com/RobertMaulana/bookstore_users_api_microservice/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNowRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNowRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError("invalid data")
	case 1292:
		return errors.NewInternalServerError("incorrect datetime value")
	}
	return errors.NewInternalServerError("error processing request")
}