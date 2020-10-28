package models

import "github.com/jinzhu/gorm"

//PatientState defines all values of a patient that can change depending on the current phase
type PatientState struct {
	gorm.Model
	RespirationRate  int `json:"respirationRate"`
	HeartRate        int `json:"heartRate"`
	OxygenSaturation int `json:"oxygenSaturation"`
	BloodPressure    int `json:"bloodPressure"`

	StateAmbulant           bool `json:"stateAmbulant"`
	StateSputteringBleeding bool `json:"stateSputteringBleeding"`
	StateBleeding           bool `json:"stateBleeding"`
	StateMotionless         bool `json:"stateMotionless"`
	StateCyanosis           bool `json:"stateCyanosis"`

	Duration           int `json:"duration"`
	NextStateA         int `gorm:"default:1"`
	NextStateB         int `gorm:"default:1"`
	NextStateC         int `gorm:"default:1"`
	ConditionPrimary   string
	ConditionSecondary string
}
