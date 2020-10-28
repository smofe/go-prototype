package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Patient defines all static properties of a patient that are independet of the current phase
type Patient struct {
	gorm.Model
	Name               string       `json:"name"`
	Age                int          `json:"age"`
	Gender             string       `json:"gender"`
	HairColor          string       `json:"hairColor"`
	PatientStateID     int          `json:"patientStateID"`
	PatientState       PatientState `gorm:"association_autoupdate:false" gorm:"association_autocreate:false" json:"patientState"` //Foreign key to current patient state
	NextPhaseTimeStamp time.Time    `json:"nextPhaseTimeStamp"`

	MeasureRecoveryPosition bool `json:"measureRecoveryPosition"`
	MeasureVentilated       bool `json:"measureVentilated"`
	MeasureTourniquet       bool `json:"measureTourniquet"`
	MeasureInfusion         bool `json:"measureInfusion"`
}
