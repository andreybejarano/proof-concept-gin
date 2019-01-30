package roles

import (
	"proof-concept-gin/db/schemas"

	"github.com/jinzhu/gorm"
)

func CreateRoles(dataBase *gorm.DB) {
	dataBase.Table("roles").CreateTable(&schemas.Rol{})
}
