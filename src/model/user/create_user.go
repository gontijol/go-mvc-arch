package model_user

import (
	logger "bv-api/src/config/logs"
	"bv-api/src/config/rest_err"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println("CreateUser", ud)

	return nil
}
