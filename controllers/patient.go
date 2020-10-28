package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smofe/go-prototype/models"
)

// CreatePatientSchema defines a valid input to create a new patient
type CreatePatientSchema struct {
	Name         string `json:"name" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	HairColor    string `json:"hairColor" binding:"required"`
	PatientState int    `json:"patientState" binding:"required"`

	MeasureRecoveryPosition bool `json:"measureRecoveryPosition"`
	MeasureVentilated       bool `json:"measureVentilated"`
	MeasureTourniquet       bool `json:"measureTourniquet"`
	MeasureInfusion         bool `json:"measureInfusion"`
}

// ReturnAllPatients GET all Patients
func ReturnAllPatients(context *gin.Context) {
	var patients []models.Patient
	models.DB.Find(&patients)
	context.JSON(http.StatusOK, patients)
}

// ReturnSinglePatient GET one single Patient
func ReturnSinglePatient(context *gin.Context) {
	var patient models.Patient
	var patientstate models.PatientState
	// Get patient with the correct id
	if err := models.DB.Where("id = ?", context.Param("id")).First(&patient).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Get the current patient state
	models.DB.Model(&patient).Related(&patientstate)
	patient.PatientState = patientstate

	context.JSON(http.StatusOK, gin.H{"data": patient})

}

// CreatePatient POST Patients
func CreatePatient(context *gin.Context) {
	var input CreatePatientSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient := models.Patient{
		Name:                    input.Name,
		Age:                     input.Age,
		Gender:                  input.Gender,
		HairColor:               input.HairColor,
		PatientStateID:          input.PatientState,
		MeasureInfusion:         input.MeasureInfusion,
		MeasureRecoveryPosition: input.MeasureRecoveryPosition,
		MeasureTourniquet:       input.MeasureTourniquet,
		MeasureVentilated:       input.MeasureVentilated,
	}
	models.DB.Create(&patient)

	context.JSON(http.StatusOK, gin.H{"data": patient})
}
