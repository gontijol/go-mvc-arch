package model_user

import (
	logger "bv-api/src/config/logs"
	"bv-api/src/config/rest_err"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) UpdateUser(id string) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "updateUser"))

	fmt.Println("UpdateUser", id)

	return nil
}
