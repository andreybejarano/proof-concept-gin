package schemas

type Rol struct {
	ID  int    `gorm:"serial PRIMARY_KEY"`
	Rol string `gorm:"text; NOT NULL"`
}
