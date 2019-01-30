package controllerusers

import (
	"net/http"
	"proof-concept-gin/db/schemas"
	usersmodel "proof-concept-gin/modules/users/model"

	"github.com/gin-gonic/gin"
)

type body struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Lastname string `form:"lastname" json:"lastname" xml:"lastname" binding:"required"`
	Mail     string `form:"mail" json:"mail" xml:"mail" binding:"required"`
	IDRol    uint   `form:"id_rol" json:"id_rol" xml:"id_rol" binding:"required"`
}

func CreateUser(context *gin.Context) {
	var body body
	context.ShouldBindJSON(&body)
	var data schemas.User
	data.Name = body.Name
	data.Lastname = body.Lastname
	data.Mail = body.Mail
	data.IDRol = body.IDRol
	response := usersmodel.CreateUser(data)
	context.JSON(http.StatusCreated, gin.H{"success": response["success"]})
}

func Users(context *gin.Context) {
	users := usersmodel.Users()
	context.JSON(http.StatusOK, gin.H{"data": users})
}
