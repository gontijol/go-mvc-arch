package model_user

import (
	"bv-api/src/config/rest_err"
	"crypto/md5"
	"encoding/hex"
)

func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
