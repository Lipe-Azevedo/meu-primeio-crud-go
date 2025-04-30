package routes

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserByID/:userid", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser/", userController.CreateUser)
	r.PUT("/updateUser/:userid", userController.UpdateUser)
	r.DELETE("/deleteUser/:userid", userController.DeleteUser)

}
