package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "fc.",
			SingularTable: false,
		}})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
