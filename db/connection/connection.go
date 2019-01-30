package connection

import (
	"fmt"
	"log"
	"proof-concept-gin/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setup() *gorm.DB {
	config.Init()
	configs := config.GetConfig()
	dbHost := fmt.Sprintf("host=%s", configs["db"]["dbHost"])
	dbPort := fmt.Sprintf(" port=%s", configs["db"]["dbPort"])
	dbName := fmt.Sprintf(" dbname=%s", configs["db"]["dbName"])
	dbUser := fmt.Sprintf(" user=%s", configs["db"]["dbUser"])
	dbPassword := fmt.Sprintf(" password=%s", configs["db"]["dbPassword"])
	db, err := gorm.Open("postgres", dbHost+dbPort+dbUser+dbName+dbPassword)
	if err != nil {
		defer log.Fatal(err)
	}
	return db
}

// Init setup inital config
func Init() *gorm.DB {
	db := setup()
	return db
}
