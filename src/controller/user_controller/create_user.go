package user_controller

import (
	"bv-api/src/config/rest_err"
	"bv-api/src/model/user/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some invalid fields=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return

	}
	fmt.Println(userRequest)
}
