package rolesmodel

import (
	"proof-concept-gin/db/connection"
	"proof-concept-gin/db/schemas"
)

func CreateRol(rol schemas.Rol) map[string]bool {
	rolToCreate := schemas.Rol{Rol: rol.Rol}
	dataBase := connection.Init()
	var response = make(map[string]bool)
	dataBase.Table("roles").Create(&rolToCreate)
	response["success"] = true
	defer dataBase.Close()
	return response
}

func Roles() []schemas.Rol {
	dataBase := connection.Init()
	roles := []schemas.Rol{}
	dataBase.Table("roles").Find(&roles)
	defer dataBase.Close()
	return roles
}
