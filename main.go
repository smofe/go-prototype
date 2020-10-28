package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/smofe/go-prototype/controllers"
	_ "github.com/smofe/go-prototype/controllers"
	"github.com/smofe/go-prototype/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func startGame(context *gin.Context) {
	// var patients []models.Patient
	// models.DB.Find(&patients)
	// for _, patient := range patients {
	// 	var patientstate models.PatientState
	// 	models.DB.Model(patient).Related(&patientstate)
	// 	_, conditionMet := models.DB.Model(patient).Get(patientstate.ConditionPrimary)
	// 	fmt.Println(conditionMet)

	// }
	// Initialize phase change times
	var patients []models.Patient
	models.DB.Find(&patients)
	for _, patient := range patients {
		var patientstate models.PatientState
		models.DB.Model(patient).Related(&patientstate)
		patient.NextPhaseTimeStamp = time.Now().Add(time.Duration(patientstate.Duration) * time.Second)
		models.DB.Save(patient)
		fmt.Println("Next Phase Change of ", patient.Name, " is at: ", patient.NextPhaseTimeStamp)
	}

	// Start asynchronous timer that checks phase changes every second
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				var patients []models.Patient
				models.DB.Find(&patients)
				for _, patient := range patients {
					//fmt.Println("current time: ", time.Now())
					//fmt.Println(patient.Name, " time: ", patient.NextPhaseTimeStamp)
					if patient.NextPhaseTimeStamp.Before(time.Now()) {
						fmt.Println(patient.Name + " just changed its state")
						var currentState models.PatientState
						models.DB.Model(patient).Related(&currentState)
						patient.PatientStateID = currentState.NextStateA
						models.DB.Save(patient)
						var nextState models.PatientState
						models.DB.Model(patient).Related(&nextState)
						patient.NextPhaseTimeStamp = time.Now().Add(time.Duration(nextState.Duration) * time.Second)
						models.DB.Save(patient)
					}
				}
			}
		}
	}()

}

func handleRequests() {
	myRouter := gin.Default()
	myRouter.POST("/patients", controllers.CreatePatient)
	myRouter.GET("/patients", controllers.ReturnAllPatients)
	myRouter.GET("/patients/:id", controllers.ReturnSinglePatient)
	myRouter.POST("/patientstates", controllers.CreatePatientState)
	myRouter.GET("/patientstates", controllers.ReturnAllPatientStates)
	myRouter.GET("/patientstates/:id", controllers.ReturnSinglePatientState)
	myRouter.GET("/startgame", startGame)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	models.ConnectDataBase()

	handleRequests()
}
