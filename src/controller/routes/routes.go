package routes

import (
	"bv-api/src/controller/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:id", user_controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", user_controller.FindUserByEmail)
	r.POST("/createUser", user_controller.CreateUser)
	r.PUT("/updateUser/:id", user_controller.UpdateUser)
	r.DELETE("/deleteUser/:id", user_controller.DeleteUser)
}
