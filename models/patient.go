package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Patient defines all static properties of a patient that are independet of the current phase
type Patient struct {
	gorm.Model
	ID                 uint         `json:"id" gorm:"primary_key"`
	Name               string       `json:"title"`
	Age                int          `json:"age"`
	Gender             string       `json:"gender"`
	HairColor          string       `json:"hairColor"`
	PatientState       PatientState `json:"patientState"` //Foreign key to current patient state
	NextPhaseTimeStamp time.Time    `json:"nextPhaseTimeStamp"`

	MeasureRecoveryPosition bool `json:"measureRecoveryPosition"`
	MeasureVentilated       bool `json:"measureVentilated"`
	MeasureTourniquet       bool `json:"measureTourniquet"`
	MeasureInfusion         bool `json:"measureInfusion"`
}
