package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	ConditionPrimary   bool `json:"conditionPrimary"`
	ConditionSecondary bool `json:"conditionSecondary"`
}

// UpdatePatientSchema defines valid input to update a new patient
type UpdatePatientSchema struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       string `json:"gender"`
	HairColor    string `json:"hairColor"`
	PatientState int    `json:"patientState"`

	MeasureRecoveryPosition *bool `json:"measureRecoveryPosition"`
	MeasureVentilated       *bool `json:"measureVentilated"`
	MeasureTourniquet       *bool `json:"measureTourniquet"`
	MeasureInfusion         *bool `json:"measureInfusion"`

	ConditionPrimary   *bool `json:"conditionPrimary"`
	ConditionSecondary *bool `json:"conditionSecondary"`
}

// ReturnAllPatients GET all Patients
func ReturnAllPatients(context *gin.Context) {
	var patients []models.Patient
	models.DB.Preload("PatientState").Find(&patients)
	context.JSON(http.StatusOK, patients)
}

// ReturnSinglePatient GET one single Patient
func ReturnSinglePatient(context *gin.Context) {
	var patient models.Patient

	if err := models.DB.Preload("PatientState").Where("id = ?", context.Param("id")).First(&patient).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

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
		MeasureInfusion:         &input.MeasureInfusion,
		MeasureRecoveryPosition: &input.MeasureRecoveryPosition,
		MeasureTourniquet:       &input.MeasureTourniquet,
		MeasureVentilated:       &input.MeasureVentilated,
		ConditionPrimary:        &input.ConditionPrimary,
		ConditionSecondary:      &input.ConditionSecondary,
	}
	models.DB.Create(&patient)

	context.JSON(http.StatusOK, gin.H{"data": patient})
}

// UpdatePatient POST Patients
func UpdatePatient(context *gin.Context) {
	// Getting the json data of the request body
	// TODO: There has to be a better way...
	var bodyBytes []byte
	if context.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(context.Request.Body)
	}
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	var bodyJSON map[string]interface{}
	err := json.Unmarshal(bodyBytes, &bodyJSON)
	if err != nil {
		fmt.Println("Error parsing request body into json object")
	}

	var patient models.Patient
	if err := models.DB.Preload("PatientState").Where("id = ?", context.Param("id")).First(&patient).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdatePatientSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var boolTrue = true
	var boolFalse = false
	for key, value := range bodyJSON {
		if value == true {
			if key == patient.PatientState.ConditionPrimary {
				fmt.Println("Primary Condition Met")
				patient.ConditionPrimary = &boolTrue
			}
			if key == patient.PatientState.ConditionSecondary {
				patient.ConditionSecondary = &boolTrue
			}
		} else {
			if key == patient.PatientState.ConditionPrimary {
				fmt.Println("Primary Condition Met")
				patient.ConditionPrimary = &boolFalse
			}
			if key == patient.PatientState.ConditionSecondary {
				patient.ConditionSecondary = &boolFalse
			}
		}
	}

	models.DB.Save(&patient)
	models.DB.Model(&patient).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": patient})
}
