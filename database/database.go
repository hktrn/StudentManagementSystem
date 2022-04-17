package database

import (
	"log"
	"os"

	"github.com/hktrn/StudentManagementSystem/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectToDatabase() {
	db, err := gorm.Open(sqlite.Open("StudentManagementSystem.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&model.Student{}, &model.Placement{}, &model.Record{})

	Database = DBInstance{DB: db}

}
//done
