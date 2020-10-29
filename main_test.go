package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHomePage(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"hello": "world",
	}
	// Grab our router
	router := handleRequests()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/", nil)
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["hello"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["hello"], value)
}

func TestCreatePatient(t *testing.T) {
	router := handleRequests()
	// body := gin.H{
	// 	"name":                    "Test Patient",
	// 	"age":                     20,
	// 	"gender":                  "male",
	// 	"hairColor":               "blonde",
	// 	"patientState":            2,
	// 	"measureRecoveryPosition": true,
	// 	"measureVentilated":       false,
	// 	"measureTourniquet":       true,
	// 	"measureInfusion":         false,
	// }
	payload := []byte(`{
		"name":                    "Test Patient",
		"age":                     20,
		"gender":                  "male",
		"hairColor":               "blonde",
		"patientState":            2
	}`)
	// 	controllers.CreatePatientSchema{
	// 		Name:                    "Test Patient",
	// 		Age:                     20,
	// 		Gender:                  "male",
	// 		HairColor:               "black",
	// 		PatientState:            1,
	// 		MeasureRecoveryPosition: true,
	// 		MeasureVentilated:       false,
	// 		MeasureTourniquet:       true,
	// 		MeasureInfusion:         false,
	// 	},
	// )
	fmt.Println(string(payload))
	// router := handleRequests()
	// w := performRequest(router, "POST", "/patients", bytes.NewReader(payload))
	// assert.Equal(t, http.StatusOK, w.Code)

	// fmt.Println(w.Body.String())

	// var response map[string]string
	// err := json.Unmarshal([]byte(w.Body.String()), &response)

	// value, exists := response["name"]

	// assert.Nil(t, err)
	// assert.True(t, exists)
	// assert.Equal(t, body["name"], value)
	req, err := http.NewRequest("POST", "/patients", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
