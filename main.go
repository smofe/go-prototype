package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/smofe/go-prototype/controllers"
	_ "github.com/smofe/go-prototype/controllers"
	"github.com/smofe/go-prototype/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func startGame(context *gin.Context) {
	//Initialize phase change times
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
	go func() {
		for {
			select {
			case <-ticker.C:
				var patients []models.Patient
				models.DB.Find(&patients)
				for _, patient := range patients {
					if patient.NextPhaseTimeStamp.Before(time.Now()) {
						fmt.Println(patient.Name + " just changed its state")
						var currentState models.PatientState
						models.DB.Model(patient).Related(&currentState)

						if *patient.ConditionPrimary {
							patient.PatientStateID = currentState.NextStateC
						} else if *patient.ConditionSecondary {
							patient.PatientStateID = currentState.NextStateB
						} else {
							patient.PatientStateID = currentState.NextStateA
						}
						models.DB.Save(patient)

						//get duration of the new state to update the next phase change timestamp of the patient
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

func homePage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func handleRequests() *gin.Engine {
	myRouter := gin.Default()
	myRouter.POST("/patients", controllers.CreatePatient)
	myRouter.GET("/patients", controllers.ReturnAllPatients)
	myRouter.GET("/patients/:id", controllers.ReturnSinglePatient)
	myRouter.PATCH("/patients/:id", controllers.UpdatePatient)
	myRouter.POST("/patientstates", controllers.CreatePatientState)
	myRouter.GET("/patientstates", controllers.ReturnAllPatientStates)
	myRouter.GET("/patientstates/:id", controllers.ReturnSinglePatientState)
	myRouter.GET("/startgame", startGame)
	myRouter.GET("/", homePage)
	return myRouter
}

func main() {
	models.ConnectDataBase()

	router := handleRequests()
	router.Run(":8000")
}
