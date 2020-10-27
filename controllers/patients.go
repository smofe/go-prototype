package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smofe/go-prototype/models"
)

func returnAllPatients(context *gin.Context) {
	var patients []models.Patient
	models.DB.Find(&patients)
	context.JSON(http.StatusOK, patients)
}
