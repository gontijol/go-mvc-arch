package user_controller

import (
	"bv-api/src/config/validation"
	"bv-api/src/model/user/request"
	"bv-api/src/model/user/response"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("CreateUser")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error tryng to marshal object, error = %s\n", err.Error())
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
}
