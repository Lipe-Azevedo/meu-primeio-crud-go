package routes

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByID/:userid", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:userid", controller.UpdateUser)
	r.DELETE("/deleteUser/:userid", controller.DeleteUser)

}
