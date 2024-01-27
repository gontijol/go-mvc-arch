package model_user

import (
	logger "bv-api/src/config/logs"
	"bv-api/src/config/rest_err"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) FindUser(id string) *rest_err.RestErr {
	logger.Info("Init findUser model", zap.String("journey", "ffindUser"))

	fmt.Println("FindUser", id)

	return nil
}
