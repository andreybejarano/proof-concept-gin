package schemas

type User struct {
	ID       int    `gorm:"serial PRIMARY_KEY"`
	Name     string `gorm:"text; NOT NULL"`
	Lastname string `gorm:"text; NOT NULL"`
	Mail     string `gorm:"text; NOT NULL"`
	IDRol    uint   `sql:"type:integer REFERENCES roles(id)"`
	Rol      []Rol  `gorm:"foreignkey:users_roles;AssociationForeignKey:IDRol"`
}
