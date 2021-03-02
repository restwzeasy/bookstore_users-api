package users

import (
	"github.com/restwzeasy/bookstore_users-api/pkg/api/errors"
)

func DoGetUser(userId int64) (*User, *errors.RestErr) {
	result := &User{
		Id: userId,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func DoCreateUser(user User) (*User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUser() {

}
