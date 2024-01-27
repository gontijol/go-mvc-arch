package user_controller

import (
	logger "bv-api/src/config/logs"
	"bv-api/src/config/validation"
	"bv-api/src/model/user/request"
	"bv-api/src/model/user/response"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Criando usuário",
		zap.String("journey", "create_user"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.InfoError("Erro ao fazer o bind do JSON", err,
			zap.String("journey", "create_user"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return

	}

	responseUser := response.UserResponse{
		ID:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	fmt.Println(responseUser)

	c.JSON(200, responseUser)
	logger.Info("User created successfuly ",
		zap.String("journey", "create_user"))

}
