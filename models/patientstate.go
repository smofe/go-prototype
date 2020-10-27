package models

//PatientState defines all values of a patient that can change depending on the current phase
type PatientState struct {
	ID               uint `json:"id" gorm:"primary_key"`
	RespirationRate  int  `json:"respirationRate"`
	HeartRate        int  `json:"heartRate"`
	OxygenSaturation int  `json:"oxygenSaturation"`
	BloodPressure    int  `json:"bloodPressure"`

	StateAmbulant           bool `json:"stateAmbulant"`
	StateSputteringBleeding bool `json:"stateSputteringBleeding"`
	StateBleeding           bool `json:"stateBleeding"`
	StateMotionless         bool `json:"stateMotionless"`
	StateCyanosis           bool `json:"stateCyanosis"`
}
