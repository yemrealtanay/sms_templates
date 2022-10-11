package main

import (
	"github.com/yemrealtanay/sms_templates/initializers"
	"github.com/yemrealtanay/sms_templates/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	if initializers.DB.Migrator().HasTable(&models.TestSmsTemplate{}) != true {
		err := initializers.DB.AutoMigrate(&models.TestSmsTemplate{})
		if err != nil {
			log.Fatal("Failed to create table")
		} else {
			log.Println("table created successfully")
		}
	} else {
		log.Println("No need to table creation")
	}

	if initializers.DB.Migrator().HasTable(&models.TestSmsTemplateType{}) != true {
		err := initializers.DB.AutoMigrate(&models.TestSmsTemplateType{})
		if err != nil {
			log.Fatal("Failed to create table")
		} else {
			log.Println("table created successfully")
		}
	} else {
		log.Println("No need to table creation")
	}

	if initializers.DB.Migrator().HasTable(&models.TestSmsTemplateCategory{}) != true {
		err := initializers.DB.AutoMigrate(&models.TestSmsTemplateCategory{})
		if err != nil {
			log.Fatal("Failed to create table")
		} else {
			log.Println("table created successfully")
		}
	} else {
		log.Println("No need to table creation")
	}

}
