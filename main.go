package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smofe/go-prototype/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database *gorm.DB

func homePage(context *gin.Context) {
	// fmt.Fprintf(writer, "Hello World")
	// fmt.Println("Endpoint Hit: homePage")
}

func returnAllPatients(context *gin.Context) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	var patients []Patient
	database.Find(&patients)
	context.JSON(http.StatusOK, patients)
}

func returnSinglePatient(context *gin.Context) {
	// vars := mux.Vars(request)
	// key := vars["id"]

	// var patient Patient
	// database.First(&patient, key)
	// json.NewEncoder(writer).Encode(patient)
}

func createNewPatient(context *gin.Context) {
	// requestBody, _ := ioutil.ReadAll(request.Body)
	// var patient Patient
	// json.Unmarshal(requestBody, &patient)
	// fmt.Println(patient)
	// database.Create(&patient)
	// json.NewEncoder(writer).Encode(patient)
}

func returnAllPatientStates(context *gin.Context) {
	// var patientstates []PatientState
	// database.Find(&patientstates)
	// for _, patientstate := range patientstates {
	// 	json.NewEncoder(writer).Encode(patientstate)
	// }
}

func returnSinglePatientState(context *gin.Context) {
	// vars := mux.Vars(request)
	// key := vars["id"]

	// var patientstate PatientState
	// database.First(&patientstate, key)
	// json.NewEncoder(writer).Encode(patientstate)
}

func createNewPatientState(context *gin.Context) {
	// requestBody, _ := ioutil.ReadAll(request.Body)
	// var patientstate PatientState
	// json.Unmarshal(requestBody, &patientstate)
	// database.Create(&patientstate)
	// json.NewEncoder(writer).Encode(patientstate)
}

func startGame(context *gin.Context) {
	// ticker := time.NewTicker(time.Second)
	// done := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case <-ticker.C:
	// 			for _, patient := range Patients {
	// 				if patient.NextPhase.Before(time.Now()) {
	// 					fmt.Println(patient.Name + " just changed its state")
	// 				}
	// 			}
	// 		}
	// 	}
	// }()

}

func resetDataBase(context *gin.Context) {
	database.Exec("DELETE FROM patients")
}

func handleRequests() {
	myRouter := gin.Default()
	myRouter.GET("/", homePage)
	myRouter.POST("/patients", createNewPatient)
	myRouter.GET("/patients", returnAllPatients)
	myRouter.GET("/patients/{id}", returnSinglePatient)
	myRouter.POST("/patientstates", createNewPatientState)
	myRouter.GET("/patientstates", returnAllPatientStates)
	myRouter.GET("/patientstates/{id}", returnSinglePatientState)
	myRouter.GET("/startgame", startGame)
	myRouter.GET("/resetdb", resetDataBase)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	// var err error
	// database, err = gorm.Open("sqlite3", "patients.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// fmt.Println("connected to db")
	// database.AutoMigrate(&Patient{})
	// database.AutoMigrate(&PatientState{})

	// defer database.Close()

	models.ConnectDatabase()

	handleRequests()
}
