package modelusers

import (
	"proof-concept-gin/db/connection"
	"proof-concept-gin/db/schemas"
)

func CreateUser(user schemas.User) map[string]bool {
	userToCreate := schemas.User{
		Name:     user.Name,
		Lastname: user.Lastname,
		Mail:     user.Mail,
		IDRol:    user.IDRol}
	dataBase := connection.Init()
	var response = make(map[string]bool)
	dataBase.Table("users").Create(&userToCreate)
	response["success"] = true
	return response
}

func Users() []schemas.User {
	dataBase := connection.Init()
	users := []schemas.User{}
	dataBase.Table("users").Find(&users)
	defer dataBase.Close()
	return users
}
