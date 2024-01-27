package model_user

import (
	logger "bv-api/src/config/logs"
	"bv-api/src/config/rest_err"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model", zap.String("journey", "deleteUser"))

	fmt.Println("DeleteUser", id)

	return nil
}
