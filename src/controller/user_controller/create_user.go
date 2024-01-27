package user_controller

import (
	"bv-api/src/config/rest_err"
	"bv-api/src/model/user/request"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error tryng to marshal object, error = %s\n", err.Error())
		errRest := rest_err.NewBadRequestError("Some fields are incorrect!")

		c.JSON(errRest.Code, errRest)
		return

	}
	fmt.Println(userRequest)
}
