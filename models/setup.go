package models

import (
	"github.com/jinzhu/gorm"
)

// DB is a pointer to the database
var DB *gorm.DB

// ConnectDataBase establishes connection to the current database
func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "simulation.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Patient{}, PatientState{})
	database.Model(&PatientState{}).AddForeignKey("patient_id", "patients(id)", "RESTRICT", "RESTRICT")

	DB = database
}
