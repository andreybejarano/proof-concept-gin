package users

import (
	"proof-concept-gin/db/schemas"

	"github.com/jinzhu/gorm"
)

func CreateUsers(dataBase *gorm.DB) {
	dataBase.Table("users").CreateTable(&schemas.User{})
}
