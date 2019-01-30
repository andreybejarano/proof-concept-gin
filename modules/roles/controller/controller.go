package rolescontroller

import (
	"net/http"
	"proof-concept-gin/db/schemas"
	rolesmodel "proof-concept-gin/modules/roles/model"

	"github.com/gin-gonic/gin"
)

type body struct {
	Rol string `form:"rol" json:"rol" xml:"rol" binding:"required"`
}

func CreateRol(context *gin.Context) {
	var body body
	context.ShouldBindJSON(&body)
	var data schemas.Rol
	data.Rol = body.Rol
	response := rolesmodel.CreateRol(data)
	context.JSON(http.StatusCreated, gin.H{"success": response["success"]})
}

func Roles(context *gin.Context) {
	roles := rolesmodel.Roles()
	context.JSON(http.StatusOK, gin.H{"data": roles})
}
