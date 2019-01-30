package main

import (
	connection "proof-concept-gin/db/connection"
	"proof-concept-gin/db/migrations/roles"
	"proof-concept-gin/db/migrations/users"
)

func main() {
	dataBase := connection.Init()
	roles.CreateRoles(dataBase)
	users.CreateUsers(dataBase)
	defer dataBase.Close()
}
