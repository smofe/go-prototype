package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smofe/go-prototype/models"
)

// CreatePatientStateSchema defines a valid input to create a new patient state
type CreatePatientStateSchema struct {
	RespirationRate  int `json:"respirationRate" binding:"required"`
	HeartRate        int `json:"heartRate" binding:"required"`
	OxygenSaturation int `json:"oxygenSaturation" binding:"required"`
	BloodPressure    int `json:"bloodPressure" binding:"required"`

	StateAmbulant           bool `json:"stateAmbulant"`
	StateSputteringBleeding bool `json:"stateSputteringBleeding"`
	StateBleeding           bool `json:"stateBleeding"`
	StateMotionless         bool `json:"stateMotionless"`
	StateCyanosis           bool `json:"stateCyanosis"`

	Duration           int    `json:"duration" binding:"required"`
	NextStateA         int    `gorm:"default:1"`
	NextStateB         int    `gorm:"default:1"`
	NextStateC         int    `gorm:"default:1"`
	ConditionPrimary   string `gorm:"default:MeasureVentilated"`
	ConditionSecondary string `gorm:"default:MeasureTourniquet"`
}

// ReturnAllPatientStates GET PatientStates
func ReturnAllPatientStates(context *gin.Context) {
	var patientstates []models.PatientState
	models.DB.Find(&patientstates)
	context.JSON(http.StatusOK, patientstates)
}

// ReturnSinglePatientState GET one single PatientState
func ReturnSinglePatientState(context *gin.Context) {
	var patientstate models.PatientState
	// Get patient state with the correct id
	if err := models.DB.Where("id = ?", context.Param("id")).First(&patientstate).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": patientstate})
}

// CreatePatientState POST PatientState
func CreatePatientState(context *gin.Context) {
	var input CreatePatientStateSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patientstate := models.PatientState{
		RespirationRate:         input.RespirationRate,
		HeartRate:               input.HeartRate,
		OxygenSaturation:        input.OxygenSaturation,
		BloodPressure:           input.BloodPressure,
		StateAmbulant:           input.StateAmbulant,
		StateBleeding:           input.StateBleeding,
		StateCyanosis:           input.StateCyanosis,
		StateMotionless:         input.StateMotionless,
		StateSputteringBleeding: input.StateSputteringBleeding,
		Duration:                input.Duration,
		NextStateA:              input.NextStateA,
		NextStateB:              input.NextStateB,
		NextStateC:              input.NextStateC,
		ConditionPrimary:        input.ConditionPrimary,
		ConditionSecondary:      input.ConditionSecondary,
	}
	models.DB.Create(&patientstate)

	context.JSON(http.StatusOK, gin.H{"data": patientstate})
}
