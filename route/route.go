package route

import (
	rolescontroller "proof-concept-gin/modules/roles/controller"
	userscontroller "proof-concept-gin/modules/users/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter module setup router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/users", userscontroller.CreateUser)
	router.GET("/users", userscontroller.Users)
	router.POST("/roles", rolescontroller.CreateRol)
	router.GET("/roles", rolescontroller.Roles)

	return router
}
