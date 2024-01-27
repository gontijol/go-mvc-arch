package user_controller

import (
	logger "bv-api/src/config/logs"
	"bv-api/src/config/validation"
	model_user "bv-api/src/model/user"
	"bv-api/src/model/user/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model_user.UserDomainInterface
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

	domainUser := model_user.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainUser.CreateUser()

	if err := domainUser.CreateUser(); err != nil {
		logger.InfoError("Erro ao criar usuário", err,
			zap.String("journey", "create_user"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("User created successfuly 🤩",
		zap.String("journey", "create_user"))

	c.JSON(200, domainUser)

}
